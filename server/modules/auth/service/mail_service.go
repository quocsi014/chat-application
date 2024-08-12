package service

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/quocsi014/common/app_error"
)


var (
	gsmtpHost string = "smtp.gmail.com"
	gsmtpPort string = "587"
	senderEmail string = "chatapp.verify@gmail.com"
)

type EmailService struct{
	auth smtp.Auth
	sender string
	smtpHost string
	smtpPort string
}
func NewGEmailService() *EmailService{

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	auth := smtp.PlainAuth("", senderEmail, os.Getenv("GSMTP_PASSWORD"), gsmtpHost)
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
