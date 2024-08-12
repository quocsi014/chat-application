package service

import (
	"log"
	"net/smtp"
	"os"
	"github.com/quocsi014/common/app_error"
)

var (
	gsmtpHost string = "smtp.gmail.com"
	gsmtpPort string = "587"
	senderEmail string = "chatapp.verify@gmail.com"
	senderPassword string = os.Getenv("GSMTP_PASSWORD")
)
type EmailService struct{
	auth smtp.Auth
	sender string
	smtpHost string
	smtpPort string
}
func NewGEmailService() *EmailService{
	auth := smtp.PlainAuth("", senderEmail, senderPassword, gsmtpHost)
	return &EmailService{
		auth: auth,
		sender: senderEmail,
		smtpHost: gsmtpHost,
		smtpPort: gsmtpPort,
	}
}

func (ms *EmailService)SendOtp(receiver ,otp string) error{
	err := smtp.SendMail(ms.smtpHost + ":" + ms.smtpPort, ms.auth, ms.sender, []string{receiver}, []byte(otp))
	if err != nil{
		log.Fatal(err)
		return app_error.ErrInternal(err)
	}
	return nil

}
