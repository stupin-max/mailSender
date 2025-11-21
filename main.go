package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mailSender/internal/file_reader"
	"mailSender/internal/mail_sender"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//вынести в конфиг файл
	settings := mail_sender.GetSettings()
	settings.SmtpHost = os.Getenv("SMTP_HOST")
	settings.SmtpPort = os.Getenv("SMTP_PORT")
	settings.From = os.Getenv("FROM")
	settings.Password = os.Getenv("PASSWORD")

	emailsPath := os.Getenv("EMAILS_PATH")
	templatesPath := os.Getenv("TEMPLATES_PATH")

	auth := smtp.PlainAuth("", settings.From, settings.Password, settings.SmtpHost)

	test, err := file_reader.FileReader(emailsPath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles(templatesPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона: %v", err)
	}
	for _, line := range test.Lines {
		fmt.Printf("Name: %v, Email: %v\n", line.Name, line.Email)
		receiver := mail_sender.Name{Name: line.Name}

		var body bytes.Buffer
		if err := tmpl.Execute(&body, receiver); err != nil {
			log.Fatalf("Ошибка применения шаблона: %v", err)
		}
		err = smtp.SendMail(
			settings.SmtpHost+":"+settings.SmtpPort,
			auth,
			settings.From,
			[]string{line.Email},
			body.Bytes(),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}
