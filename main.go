package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mailSender/internal/config"
	"mailSender/internal/file_reader"
	"mailSender/internal/mail_sender"
	"net/smtp"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	smptSetting := config.GetSmtpSettings()
	mailSettings := config.GetMailSettings()

	auth := smtp.PlainAuth("", smptSetting.From, smptSetting.Password, smptSetting.SmtpHost)
	test, err := file_reader.FileReader(mailSettings.EmailsPath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles(mailSettings.TemplatesPath)
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
			smptSetting.SmtpHost+":"+smptSetting.SmtpPort,
			auth,
			smptSetting.From,
			[]string{line.Email},
			body.Bytes(),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}
