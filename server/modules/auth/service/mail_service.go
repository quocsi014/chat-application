package service

import (
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/quocsi014/common/app_error"
)

var (
	gsmtpHost string = "smtp.gmail.com"
	gsmtpPort string = "587"
	senderEmail string = "chatapp.verify@gmail.com"
	senderPassword string = os.Getenv("GSMTP_PASSWORD")
)
type MailService struct{
	auth smtp.Auth
	sender string
	smtpHost string
	smtpPort string
}
func NewGMailService() *MailService{
	auth := smtp.PlainAuth("", senderEmail, senderPassword, gsmtpHost)
	return &MailService{
		auth: auth,
		sender: senderEmail,
		smtpHost: gsmtpHost,
		smtpPort: gsmtpPort,
	}
}

func (ms *MailService)SendOtp(receiver ,otp string) error{
	err := smtp.SendMail(ms.smtpHost + ":" + ms.smtpPort, ms.auth, ms.sender, []string{receiver}, []byte(otp))
	if err != nil{
		log.Fatal(err)
		return app_error.ErrInternal(err)
	}
	return nil

}
