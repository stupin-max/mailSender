package app

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mailSender/internal/config"
	"mailSender/internal/file_reader"
	"mailSender/internal/mail_sender"
	"net/smtp"
	"sync"

	"github.com/joho/godotenv"
)

// App represents the application structure
type App struct {
	smtpSettings *config.SmtpSetting
	mailSettings *config.MailSettings
	tmpl         *template.Template
	auth         smtp.Auth
	address      string
}

// New creates and initializes a new App instance
func New() (*App, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	smtpSettings, err := config.GetSmtpSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to get SMTP settings: %w", err)
	}

	mailSettings, err := config.GetMailSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to get mail settings: %w", err)
	}

	tmpl, err := template.ParseFiles(mailSettings.TemplatesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	auth := smtp.PlainAuth("", smtpSettings.From, smtpSettings.Password, smtpSettings.SmtpHost)
	address := fmt.Sprintf("%s:%s", smtpSettings.SmtpHost, smtpSettings.SmtpPort)

	return &App{
		smtpSettings: smtpSettings,
		mailSettings: mailSettings,
		tmpl:         tmpl,
		auth:         auth,
		address:      address,
	}, nil
}

// Run executes the main application logic
func (a *App) Run() error {
	recipients, err := file_reader.ReadCSV(a.mailSettings.EmailsPath)
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	// Send emails concurrently using worker pool
	const numWorkers = 5 // Number of concurrent workers
	if err := a.sendEmailsConcurrently(recipients.Lines, numWorkers); err != nil {
		return fmt.Errorf("failed to send emails: %w", err)
	}

	return nil
}

func (a *App) sendEmailsConcurrently(recipients []file_reader.CSVLine, numWorkers int) error {
	if len(recipients) == 0 {
		return nil
	}

	// Create a channel to send work items
	jobs := make(chan file_reader.CSVLine, len(recipients))

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for recipient := range jobs {
				data := mail_sender.TemplateData{Name: recipient.Name}
				if err := a.sendEmail(recipient, data); err != nil {
					log.Printf("[Worker %d] Failed to send email to %s (%s): %v", workerID, recipient.Name, recipient.Email, err)
				} else {
					log.Printf("[Worker %d] Successfully sent email to %s (%s)", workerID, recipient.Name, recipient.Email)
				}
			}
		}(i)
	}

	// Send all recipients to the jobs channel
	for _, recipient := range recipients {
		jobs <- recipient
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	return nil
}

func (a *App) sendEmail(recipient file_reader.CSVLine, data mail_sender.TemplateData) error {
	var body bytes.Buffer
	if err := a.tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if err := smtp.SendMail(a.address, a.auth, a.smtpSettings.From, []string{recipient.Email}, body.Bytes()); err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}

	return nil
}

