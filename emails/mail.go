package emails

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var host = os.Getenv("SMTP_HOST")
var port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
var from = os.Getenv("EMAIL")
var password = os.Getenv("EMAIL_PASSWORD")

var d = gomail.NewDialer(host, port, from, password)

func SendOTP(OTP int, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP")
	m.SetBody("text/html", fmt.Sprintf("Your OTP is %v", OTP))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
