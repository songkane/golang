// Package http authentication and authorization
// Created by chenguolin 2019-02-11
package http

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-common/rsa"
)

// VerifyAuthorization verify authorization
// always return true mean verify pass, if return false
// mean verify failed
func VerifyAuthorization(c *gin.Context) error {
	// 1. get public key and signature
	pubk := c.GetString("publicKey")
	sig := c.GetString("signature")

	// 2. sort form values
	// content-type muse be application/x-www-form-urlencoded
	params := make([]string, 0, 10)
	form := c.Request.Form
	for k, v := range form {
		// filter sig field
		if k == sigField {
			continue
		}
		params = append(params, v[0])
	}
	sort.Strings(params)
	data := strings.Join(params, "")

	// 3. verify signature
	return rsa.VerifySignatureByPubKey(rsa.PublicKey(pubk), rsa.Signature(sig), []byte(data))
}
