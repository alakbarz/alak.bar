package main

import (
	"log"
	"net/smtp"
	"os"

	"gopkg.in/macaron.v1"
)

func sendMail(subject, message string, ctx *macaron.Context) {
	from := "hello@alak.bar"
	pass := os.Getenv("GPSSWRD")
	to := "report@alak.bar"
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: " + subject + "\n\n" + message

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("[Gmail] SMTP ERR: %s", err)
		ctx.Redirect("/emailerror")
		return
	}

	log.Print("[Gmail] Email sent")
	ctx.Redirect("/thankyou")
}
