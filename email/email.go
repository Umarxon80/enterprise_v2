package email

import (
	"bytes"

	"html/template"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v3/log"
)

type emailData struct {
	Name    string
	Message string
}
type otp struct {
	Name string
	Message string
	OtpLink string
}

func SendMail(to, name, subject, content string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// to := []string{"umarhonsultanov53@gmail.com"}
	data := emailData{
		Name:    name,
		Message: content,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Error("Error parsing html", err)
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Error("Error executing html", err)
		return err
	}
	subject = "Subject: " + subject + "\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	message := []byte(subject + mime + body.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Error("Error sending message", err)
		return err
	}
	return nil
}

func SendOtp(to, name, subject,content, link string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	data := otp{
		Name:    name,
		Message: content,
		OtpLink: link,
	}

	tmpl, err := template.ParseFiles("otp-template.html")
	if err != nil {
		log.Error("Error parsing html", err)
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Error("Error executing html", err)
		return err
	}
	subject = "Subject: " + subject + "\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	message := []byte(subject + mime + body.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Error("Error sending message", err)
		return err
	}
	return nil
}
