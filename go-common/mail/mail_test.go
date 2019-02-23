// Package mail unit test
// Created by chenguolin 2019-02-23
package mail

import (
	"fmt"
	"testing"
)

func TestNewMail(t *testing.T) {
	user := "xuankulgc@gmail.com"
	password := "12345xuanku"
	smtpHost := "smtp.gmail.com"
	port := 587

	mail := NewMail(user, password, smtpHost, port)
	if mail == nil {
		t.Fatal("TestNewMail mail == nil")
	}
}

func TestMail_Send(t *testing.T) {
	user := "xuankulgc@gmail.com"
	password := "12345xuanku"
	smtpHost := "smtp.gmail.com"
	port := 587

	mail := NewMail(user, password, smtpHost, port)
	if mail == nil {
		t.Fatal("TestNewMail mail == nil")
	}

	// case 1
	err := mail.Send(nil, "xuanku 2 cgl", "xuanku test data")
	fmt.Println(err)
	if err == nil {
		t.Fatal("TestMail_Send case 1 err == nil")
	}

	// case 2
	receiver := make([]string, 0)
	receiver = append(receiver, "cgl1079743846@gmail.com")
	err = mail.Send(receiver, "xuanku 2 cgl", "xuanku test send mail data")
	fmt.Println(err)
	if err != nil {
		t.Fatal("TestMail_Send case 2 err != nil")
	}
}
