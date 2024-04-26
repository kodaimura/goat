package mail

import (
	"fmt"
	"strings"
    "net/smtp"
	"mime"
	"goat/config"
)


func Send (from string, to []string, subject, body string) {
	cf := config.GetConfig()
	auth := smtp.PlainAuth("", cf.MailUser, cf.MailPass, cf.MailHost)

	header := "From: " + from + "\r\n" +
		"To: " + strings.Join(to, ", ") + "\r\n" +
		"Subject: " + mime.QEncoding.Encode("UTF-8", subject) + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n"

	mail := header + "\r\n" + body

    if err := smtp.SendMail(fmt.Sprintf("%s:%s", cf.MailHost, cf.MailPort), auth, from, to, []byte(mail)); err != nil {
        fmt.Println("Error sending email:", err)
    }
}