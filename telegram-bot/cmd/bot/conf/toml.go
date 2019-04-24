// Package config read toml file
// Created by chenguolin 2019-04-24
package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Load from file
func Load(filePath string, dest interface{}) {
	_, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile(filePath, dest)
	if err != nil {
		panic(err)
	}
}
