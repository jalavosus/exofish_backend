package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

const (
	username   string = "projects@aliancetek.com"
	smtpServer string = "smtp.mail.us-east-1.awsapps.com"
	smtpPort   int    = 465
)

func SendMail(recipientName, recipientEmail string, htmlBody, textBody string) {
	m := gomail.NewMessage()

	toAddr := fmt.Sprintf("%s <%s>", recipientName, recipientEmail)

	m.SetHeader("From", "Yushuttles Demo <projects@aliancetek.com>")
	m.SetHeader("To", toAddr)
	m.SetHeader("Subject", "Thank you for your appointment")
	m.SetBody("text/plain", textBody)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(smtpServer, smtpPort, username, os.Getenv("emailPassword"))

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
