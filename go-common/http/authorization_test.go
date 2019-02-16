// Package http signature unit test
// Created by chenguolin 2018-11-17
package http

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-common/rsa"
)

const (
	testPrivateKey = "FHOzhQe9rk0x+uPRK72FpVDl/N1eGHTBbq4VjzXto/o="
	testPublicKey  = "2x1+3Jye4DDsmHbte3YmxVR0h0QH3GIhp5OR2ZMr4Uo=pI3zPrj4CTViWgYmhZkmwO832BYFuvdv1enE54joxOs="
)

func TestVerifyAuthorization(t *testing.T) {
	// gen signature
	i := 0
	body := ""
	for i < 10000000 {
		body = string(strconv.Itoa(i))
		i++
	}
	sig, err := rsa.GenSignatureByPriKey(testPrivateKey, []byte(body))
	if err != nil {
		t.Fatal("TestVerifyAuthorization GenSignatureByPriKey err != nil")
	}

	c := &gin.Context{}
	c.Request = &http.Request{}
	form := make(map[string][]string)
	form["id"] = []string{"123456"}
	form["name"] = []string{"chenguolin"}
	form["sigTime"] = []string{"1234567890"}
	c.Request.Form = form

	c.Set("publicKey", testPublicKey)
	c.Set("signature", string(sig))

	err = VerifyAuthorization(c)
	if err != nil {
		t.Fatal("TestVerifyAuthorization VerifyAuthorization err != nil")
	}
}
