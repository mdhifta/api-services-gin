package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendEmail(subject string, message string, to []string) {
	// Connect to the SMTP Server
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD_EMAIL"), os.Getenv("HOST_EMAIL"))

	// set message
	msg := "Subject: " + subject + "\n" + message
	err := smtp.SendMail(os.Getenv("SERVERNAME_EMAIL"), auth, os.Getenv("EMAIL"), to, []byte(msg))

	if err != nil {
		fmt.Println(err)
	}
}

func SendEmailHTML(subject string, templatePath string, link string, to []string) {
	// Connect to the SMTP Server
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD_EMAIL"), os.Getenv("HOST_EMAIL"))

	// get html
	var body bytes.Buffer
	html, err := template.ParseFiles(templatePath)

	html.Execute(&body, struct{ Link string }{Link: link})

	// set message
	headers := "MIME-version: 1.0;\nContent-Type: text/html;charset=\"UTF-8\";"
	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(os.Getenv("SERVERNAME_EMAIL"), auth, os.Getenv("EMAIL"), to, []byte(msg))

	if err != nil {
		fmt.Println(err)
	}
}
