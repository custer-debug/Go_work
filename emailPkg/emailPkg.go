package emailPkg

import (
	"crypto/rand"
	"custer-debug/serverConst"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func randomCode() string {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	s := fmt.Sprintf("%x", b)
	s = strings.ToUpper(s)
	return s
}

func SendEmailCreateUser(login string) error {

	from := serverConst.GmailServer
	pass := serverConst.GmailPassword
	to := login
	serverConst.Code = randomCode()
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Registration\n\n" +
		"Thank you for registration on our site!\n" +
		"Your personal code: " + serverConst.Code +
		"\nWith respect, Packy team!"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	fmt.Print("Registration: " + to)

	return nil
}
