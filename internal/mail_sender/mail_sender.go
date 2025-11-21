package mail_sender

type Setting struct {
	SmtpHost string
	SmtpPort string
	From     string
	Password string
}

type MailTemplate struct {
	Subject string
	Body    string
}

func GetSettings() *Setting {
	return &Setting{}
}

type Name struct {
	Name string
}

func (mt *MailTemplate) SendMail() []byte {
	return []byte(mt.Subject + "\n" + mt.Body)
}

type sender struct {
	Setting      Setting
	MailTemplate MailTemplate
}
