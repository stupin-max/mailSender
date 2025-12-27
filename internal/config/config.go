package config

import (
	"fmt"
	"os"
)

type SmtpSetting struct {
	SmtpHost string
	SmtpPort string
	From     string
	Password string
}

type MailSettings struct {
	EmailsPath    string
	TemplatesPath string
}

func GetSmtpSettings() (*SmtpSetting, error) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	from := os.Getenv("FROM")
	password := os.Getenv("PASSWORD")

	if smtpHost == "" {
		return nil, fmt.Errorf("SMTP_HOST environment variable is not set")
	}
	if smtpPort == "" {
		return nil, fmt.Errorf("SMTP_PORT environment variable is not set")
	}
	if from == "" {
		return nil, fmt.Errorf("FROM environment variable is not set")
	}
	if password == "" {
		return nil, fmt.Errorf("PASSWORD environment variable is not set")
	}

	return &SmtpSetting{
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
		From:     from,
		Password: password,
	}, nil
}

func GetMailSettings() (*MailSettings, error) {
	emailsPath := os.Getenv("EMAILS_PATH")
	templatesPath := os.Getenv("TEMPLATES_PATH")

	if emailsPath == "" {
		return nil, fmt.Errorf("EMAILS_PATH environment variable is not set")
	}
	if templatesPath == "" {
		return nil, fmt.Errorf("TEMPLATES_PATH environment variable is not set")
	}

	return &MailSettings{
		EmailsPath:    emailsPath,
		TemplatesPath: templatesPath,
	}, nil
}
