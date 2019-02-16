// Package rsa unit test
// Created by chenguolin 2019-02-16
package rsa

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	testPrivateKey = "FHOzhQe9rk0x+uPRK72FpVDl/N1eGHTBbq4VjzXto/o="
	testPublicKey  = "2x1+3Jye4DDsmHbte3YmxVR0h0QH3GIhp5OR2ZMr4Uo=pI3zPrj4CTViWgYmhZkmwO832BYFuvdv1enE54joxOs="
)

func TestGenKeyPair(t *testing.T) {
	prv, pub, err := GenKeyPair()
	if err != nil {
		t.Fatal("TestGenkeyPair failed err != nil")
	}
	if len(prv) != 44 {
		t.Fatal("TestGenkeyPair len(prv) != 44 failed ~")
	}
	if len(pub) != 44*2 {
		t.Fatal("TestGenkeyPair len(pub) != 44 failed ~")
	}
	fmt.Println(prv, len(prv))
	fmt.Println(pub, len(pub))
}

func TestGenSignatureByPriKey(t *testing.T) {
	// case 1
	i := 0
	body := ""
	for i < 10000000 {
		body = string(strconv.Itoa(i))
		i++
	}
	sig, err := GenSignatureByPriKey(testPrivateKey, []byte(body))
	if err != nil {
		t.Fatal("TestGenSignatureByPriKey GenSignatureByPriKey err != nil")
	}
	if len(sig) != 208 {
		t.Fatal("TestGenSignatureByPriKey len(sig) != 208")
	}
	fmt.Println(sig)
}

func TestVerifySignatureByPubKey(t *testing.T) {
	// gen sig
	i := 0
	body := ""
	for i < 10000000 {
		body = string(strconv.Itoa(i))
		i++
	}
	sig, err := GenSignatureByPriKey(testPrivateKey, []byte(body))
	if err != nil {
		t.Fatal("TestGenSignatureByPriKey GenSignatureByPriKey err != nil")
	}

	// verify
	err = VerifySignatureByPubKey(testPublicKey, sig, []byte(body))
	fmt.Println(err)
}
