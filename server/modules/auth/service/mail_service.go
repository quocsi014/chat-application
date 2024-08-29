package service

import (
	"fmt"
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

func (ms *EmailService)SendRegistrationVerification(receiver, token string) error{
	clientVerificationUrl := os.Getenv("CLIENT_VERIFICATION_URL")
	subject := "Subject: Please Verify Your Email\r\n"
	contentType := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	body := fmt.Sprintf(`
	<html>
	<body>
		<p>Hello,</p>
		<p>You have just registered the application, please press the button below to verify your email.</p>
		<p>if it is not you please ignore this email.</p>
		<p>(Valid for 5 minutes)</p>
		<a href="%s?token=%s" style="
			display: inline-block;
			padding: 10px 20px;
			font-size: 16px;
			color: #ffffff;
			background-color: #007bff;
			text-decoration: none;
			border-radius: 5px;
		">Verify Email</a>
		<p>Thank you!</p>
	</body>
	</html>`, clientVerificationUrl, token)

	// Full email message.
	message := []byte(subject + contentType + body)
	err := smtp.SendMail(ms.smtpHost + ":" + ms.smtpPort, ms.auth, ms.sender, []string{receiver}, []byte(message))
	if err != nil{
		log.Fatal(err)
		return app_error.ErrInternal(err)
	}
	return nil
}
