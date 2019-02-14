// Package ulid ulid生成基础库
// Created by chenguolin 2018-11-16
// https://github.com/oklog/ulid
package uuid

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

// Use pool to avoid concurrent access for rand.Source
var entropyPool = sync.Pool{
	New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

// NewUlid Generate Unique ID
// Currently using ULID, this maybe conflict with other process with very low possibility
func NewUlid() string {
	entropy := entropyPool.Get().(*rand.Rand)
	defer entropyPool.Put(entropy)

	id := ulid.MustNew(ulid.Now(), entropy)
	return strings.ToLower(id.String())
}
