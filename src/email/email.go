package email

import (
	"crypto/tls"
	"fmt"

	"github.com/fatih/color"
	"gopkg.in/gomail.v2"
)

func SendEmail(host string) {
	fmt.Println("Enter to data...")
	fromFakeEmail := readFromFakeEmail()
	fromEmail, fromPassword, fromSmtp, fromPort, toEmail := readFromData()

	makeEmail(fromFakeEmail, fromEmail, fromPassword, fromSmtp, fromPort, toEmail)
}

func readFromFakeEmail() string {
	var fromFakeEmail string
	color.Yellow("\nEnter to fake email to test:")
	fmt.Scan(&fromFakeEmail)

	return fromFakeEmail
}

func readFromData() (string, string, string, int, string) {
	var fromEmail, fromPassword, fromSmtp, toEmail string
	var fromPort int

	color.Yellow("\nEnter to from email:")
	fmt.Scan(&fromEmail)

	color.Yellow("\nEnter to from password:")
	fmt.Scan(&fromPassword)

	color.Yellow("\nEnter to from smtp server:")
	fmt.Scan(&fromSmtp)

	color.Yellow("\nEnter to from port:")
	fmt.Scan(&fromPort)

	color.Yellow("\nEnter to email:")
	fmt.Scan(&toEmail)

	return fromEmail, fromPassword, fromSmtp, fromPort, toEmail
}

func makeEmail(fromFakeEmail string, fromEmail string, fromPassword string, fromSmtp string, fromPort int, toEmail string) {
	fmt.Println("Sending email...")
	mail := gomail.NewMessage()

	// Set FAKE E-Mail sender
	mail.SetHeader("From", fromFakeEmail)

	// Set E-Mail receivers
	mail.SetHeader("To", toEmail)

	mail.SetHeader("Subject", "Email send with check email spoof")
	mail.SetBody("text/plain", "Email send with check email spoof - Host vulnerability")

	d := gomail.NewDialer(fromSmtp, fromPort, fromEmail, fromPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}
	color.Green("\n Email Sent Successfully! \n")
}
