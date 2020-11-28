package emailPkg

import (
	"custer-debug/serverConst"
	"fmt"
	"log"
	"net/smtp"
)







func SendEmailCreateUser(login string)error{

	from := serverConst.GmailServer
	pass := serverConst.GmailPassword
	to := login

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Registration\n\n" +
		"Thank you for registration on our site!\n" +
		"With respect, Packy team!"


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
