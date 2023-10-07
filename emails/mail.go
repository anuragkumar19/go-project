package emails

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func SendOTP(OTP int, email string) error {

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	to := []string{
		email,
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("emails/otp.html")

	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		OTP int
	}{
		OTP: OTP,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

	if err != nil {
		return err
	}

	return nil
}
