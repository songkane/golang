// Package dlock 单元测试
// Created by chenguolin 2018-11-16
package dlock

import (
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestMemcacheDLocker_Lock(t *testing.T) {
	mcClient := memcache.Client{}
	mcLocker := NewMemcachedDLocker(mcClient)

	// case 1
	key := "test_lock_key"
	expiration := int32(10) //秒
	opt := LockRetryOptions{}
	succ, err := mcLocker.Lock(key, expiration, opt)
	if succ == false {
		t.Fatal("TestMemcachedDLocker_Lock case 1 return false")
	}
	if err != nil {
		t.Fatal("TestMemcachedDLocker_Lock case 1 error", err)
	}

	// case 2
	key = "test_lock_key2"
	expiration = int32(10) //秒
	opt = LockRetryOptions{
		Count:    3,
		Interval: time.Duration(10),
	}
	succ, err = mcLocker.Lock(key, expiration, opt)
	if succ == false {
		t.Fatal("TestMemcachedDLocker_Lock case 2 return false")
	}
	if err != nil {
		t.Fatal("TestMemcachedDLocker_Lock case 2 error", err)
	}
}

func TestMemcacheDLocker_Unlock(t *testing.T) {
	mcClient := memcache.Client{}
	mcLocker := NewMemcachedDLocker(mcClient)

	// case 1
	key := "test_lock_key"
	opt := LockRetryOptions{}
	err := mcLocker.Unlock(key, opt)
	if err != nil {
		t.Fatal("TestMemcachedDLocker_Unlock case 1 error", err)
	}

	// case 2
	key = "test_lock_key2"
	opt = LockRetryOptions{
		Count:    3,
		Interval: time.Duration(10),
	}
	err = mcLocker.Unlock(key, opt)
	if err != nil {
		t.Fatal("TestMemcachedDLocker_Unlock case 2 error", err)
	}

	// case 3
	key = "test_lock_key3"
	opt = LockRetryOptions{
		Count:    3,
		Interval: time.Duration(10),
	}
	err = mcLocker.Unlock(key, opt)
	if err != nil {
		t.Fatal("TestMemcachedDLocker_Unlock case 3 error", err)
	}
}
