// Package ticker unit test
// Created by chenguolin 2019-02-11
package ticker

import (
	"testing"
	"fmt"
	"time"
)

func printMillisecond() {
	fmt.Println("printMillisecond ...")
}

func printSecond() {
	fmt.Println("printSecond ...")
}

func printMinute() {
	fmt.Println("printMinute ...")
}

func TestTicker(t *testing.T) {
	// case 1
	Ticker(time.Duration(100 * time.Millisecond), printMillisecond)

	// case 2
	Ticker(time.Duration(2 * time.Second), printSecond)

	// case 3
	Ticker(time.Duration(1 * time.Minute), printMinute)

	time.Sleep(time.Duration(2 * time.Minute))
}
