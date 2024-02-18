package infrastructure

import (
	"crypto/tls"
	"email-api/ports/config"

	"gopkg.in/mail.v2"
)

var Mailer *mail.Dialer

func InitMailSender() {
	Mailer = mail.NewDialer(
		config.Config.MailSender.SMTP,
		config.Config.MailSender.Port,
		config.Config.MailSender.From,
		config.Config.MailSender.Password,
	)
	Mailer.TLSConfig = &tls.Config{
		InsecureSkipVerify: config.Config.MailSender.InsecureSkipVerify,
	}
}
