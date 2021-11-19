package main

import (
	"net/http"

	"gopkg.in/gomail.v2"
)

func email(w http.ResponseWriter, r *http.Request) {

	m := gomail.NewMessage()
	m.SetHeader("From", "parvezkhan12799@gmail.com")
	m.SetHeader("To", "pkkhan120799@gmail.com")
	m.SetHeader("Subject", "Hello!")
	txt := `Hello <b>Bob</b>!
	<h2>jhgjhhjdf</h2>
	<h1>hi</h1>`
	m.SetBody("text/html", txt)

	// Send the email to Bob
	d := gomail.NewDialer("smtp.mailtrap.io", 587, "c8f5171290d36a", "e7605b81c1c348")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
