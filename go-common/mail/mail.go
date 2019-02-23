// Package mail send email
// Created by chenguolin 2019-02-23
package mail

import (
	"fmt"
	"net/smtp"
)

// Mail client
type Mail struct {
	auth     smtp.Auth
	user     string
	port     int
	smtpAddr string
}

// NewMail new mail client
func NewMail(user, password, smtpHost string, port int) *Mail {
	auth := smtp.PlainAuth("", user, password, smtpHost)
	return &Mail{
		auth:     auth,
		user:     user,
		port:     port,
		smtpAddr: fmt.Sprintf("%s:%d", smtpHost, port),
	}
}

func (m *Mail) Send(receiver []string, subject, data string) error {
	if receiver == nil {
		return fmt.Errorf("receiver is nil")
	}

	// var re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// validate receiver
	// if !re.MatchString(receiver) {
	// return fmt.Errorf("validate email address failed receiver: %s", receiver)
	// }

	// format message
	msg := fmt.Sprintf("Subject: %s\r\n\r\n%s\r\n", subject, data)

	// send mail
	err := smtp.SendMail(
		m.smtpAddr,
		m.auth,
		m.user,
		receiver,
		[]byte(msg),
	)

	return err
}
