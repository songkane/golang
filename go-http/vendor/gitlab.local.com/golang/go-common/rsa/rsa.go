// Package rsa generate verify
// Created by chenguolin 2019-02-16
package rsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
)

type (
	// PrivateKey string
	PrivateKey string
	// PublicKey string
	PublicKey string
	// Signature string
	Signature string
)

// GenKeyPair generate private, public key
// - PrivateKey private key
// - PublicKey public key
func GenKeyPair() (PrivateKey, PublicKey, error) {
	var prk *ecdsa.PrivateKey
	var puk ecdsa.PublicKey
	var curve elliptic.Curve

	curve = elliptic.P256()
	prk, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return PrivateKey(""), PublicKey(""), err
	}
	puk = prk.PublicKey

	// base64 encode
	prvKey := PrivateKey(base64.StdEncoding.EncodeToString(prk.D.Bytes()))
	pubXKey := PublicKey(base64.StdEncoding.EncodeToString(puk.X.Bytes()))
	pubYKey := PublicKey(base64.StdEncoding.EncodeToString(puk.Y.Bytes()))

	return prvKey, pubXKey + pubYKey, nil
}

// GenSignatureByPriKey gen signature by private key
// @prvKey private key
// @data content data
//
// - Signature signature value
func GenSignatureByPriKey(prvKey PrivateKey, data []byte) (Signature, error) {
	// base64 decode
	prvBytes, err := base64.StdEncoding.DecodeString(string(prvKey))
	if err != nil {
		return Signature(""), err
	}

	// new ecdsa.PublicKey
	pubk := ecdsa.PublicKey{
		Curve: elliptic.P256(),
	}
	prvk := &ecdsa.PrivateKey{
		PublicKey: pubk,
		D:         fromBase10(string(prvBytes)),
	}

	// sign
	r, s, err := ecdsa.Sign(rand.Reader, prvk, data)
	if err != nil {
		return Signature(""), err
	}

	rt, err := r.MarshalText()
	if err != nil {
		return Signature(""), err
	}
	st, err := s.MarshalText()
	if err != nil {
		return Signature(""), err
	}

	// base64 encode
	rtBase64 := base64.StdEncoding.EncodeToString(rt)
	stBase64 := base64.StdEncoding.EncodeToString(st)

	return Signature(rtBase64 + stBase64), nil
}

// VerifySignatureByPubKey verify signature by public key
// @pub public key
// @sig signature
// @data content data
func VerifySignatureByPubKey(pub PublicKey, sig Signature, data []byte) error {
	// pub = pubx + puby
	if len(pub) != 88 {
		return errors.New("VerifySignatureByPubKey invalid public key")
	}
	pubx, err := base64.StdEncoding.DecodeString(string(pub[:44]))
	if err != nil {
		return err
	}
	puby, err := base64.StdEncoding.DecodeString(string(pub[44:]))
	if err != nil {
		return err
	}

	// new ecdsa.PrivateKey
	pubk := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     fromBase10(string(pubx)),
		Y:     fromBase10(string(puby)),
	}

	// parse sig get r, s
	if len(sig) != 208 {
		return errors.New("VerifySignatureByPubKey invalid signature")
	}
	r, err := base64.StdEncoding.DecodeString(string(sig[:104]))
	if err != nil {
		return err
	}
	s, err := base64.StdEncoding.DecodeString(string(sig[104:]))
	if err != nil {
		return err
	}

	// verify signature
	if ecdsa.Verify(&pubk, data, fromBase10(string(r)), fromBase10(string(s))) {
		return errors.New("VerifySignatureByPubKey ecdsa.Verify failed ~")
	}

	return nil
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		// TODO print error log
		// TODO default return 0
		return big.NewInt(0)
	}

	return i
}
