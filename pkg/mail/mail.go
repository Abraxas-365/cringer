package mail

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/Abraxas-365/cringer/internal"
)

type Mail struct {
	From     string
	Password string
	To       []string
	Msg      string
}

func (m *Mail) SendMail(duration time.Time) error {

	elapsed := time.Since(duration)

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(m.Msg + " " + "Duration: " + internal.ShortDur(elapsed))

	// Authentication.
	auth := smtp.PlainAuth("", m.From, m.Password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, m.From, m.To, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
