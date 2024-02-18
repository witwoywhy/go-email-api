package sendout

import (
	"email-api/ports/config"

	"gopkg.in/mail.v2"
)

type adaptorMail struct {
	mail *mail.Dialer
}

func NewAdaptorMail(mail *mail.Dialer) Port {
	return &adaptorMail{
		mail: mail,
	}
}

func (a *adaptorMail) Execute(request Request) error {
	m := mail.NewMessage()
	m.SetHeader("From", config.Config.MailSender.From)
	m.SetHeader("To", request.To)
	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/html", request.Body)

	for _, v := range request.Attachments {
		m.AttachReader(v.FileName, v.Object)
	}

	return a.mail.DialAndSend(m)
}
