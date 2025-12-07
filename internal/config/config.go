package config

import "os"

type SmptSetting struct {
	SmtpHost string
	SmtpPort string
	From     string
	Password string
}

type MailSettings struct {
	EmailsPath    string
	TemplatesPath string
}

func GetSmtpSettings() *SmptSetting {
	return &SmptSetting{
		SmtpHost: os.Getenv("SMTP_HOST"),
		SmtpPort: os.Getenv("SMTP_PORT"),
		From:     os.Getenv("FROM"),
		Password: os.Getenv("PASSWORD"),
	}
}

func GetMailSettings() *MailSettings {
	return &MailSettings{
		EmailsPath:    os.Getenv("EMAILS_PATH"),
		TemplatesPath: os.Getenv("TEMPLATES_PATH"),
	}
}
