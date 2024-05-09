package mail

import (
	"fmt"
	"go.uber.org/zap"
	"net/smtp"
)

type MailServer struct {
	Email    string
	Password string
}

func NewMailServer(email, password string) *MailServer {
	return &MailServer{
		Email:    email,
		Password: password,
	}
}

func (m *MailServer) Send(to, message string) error {
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", m.Email, m.Password, smtpHost)

	// Sending email.
	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, m.Email, []string{to}, []byte("Subject: Your OTP Code\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
		"\r\n"+
		message),
	)
	if err != nil {
		zap.L().Error("Error while sending email", zap.Error(err))
		return err
	}

	return nil
}
