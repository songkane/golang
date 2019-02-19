// Package lrucache unit test
// Created by chenguolin 2019-02-18
package lrucache

import (
	"fmt"
	"testing"
)

func TestNewLruCache(t *testing.T) {
	cache := NewLruCache(10)
	if cache.capacity != 10 {
		t.Fatal("TestNewLruCache cache.capacity != 10")
	}
	if cache.size != 0 {
		t.Fatal("TestNewLruCache cache.size != 0")
	}
	if len(cache.hMap) != 0 {
		t.Fatal("TestNewLruCache cache.hMap len != 0")
	}
}

func TestLruCache_Set(t *testing.T) {
	cache := NewLruCache(3)

	// case 1 insert 1
	cache.Set("key1", 1)
	dlStr := cache.dl.print()
	if dlStr != "key1" {
		t.Fatal("case 1 dlStr != \"key1\"")
	}

	// case 2 insert 2
	cache.Set("key2", 2)
	dlStr = cache.dl.print()
	if dlStr != "key2 -> key1" {
		t.Fatal("case 2 dlStr != \"key2 -> key1\"")
	}

	// case 3 insert 3
	cache.Set("key3", 3)
	dlStr = cache.dl.print()
	if dlStr != "key3 -> key2 -> key1" {
		t.Fatal("case 3 dlStr != \"key3 -> key2 -> key1\"")
	}

	// case 4 insert 4
	cache.Set("key4", 4)
	dlStr = cache.dl.print()
	if dlStr != "key4 -> key3 -> key2" {
		t.Fatal("case 4 dlStr != \"key4 -> key3 -> key2\"")
	}

	// case 5 insert 1
	cache.Set("key1", 1)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key1 -> key4 -> key3" {
		t.Fatal("case 5 dlStr != \"key1 -> key4 -> key3\"")
	}

	// case 6 insert 2
	cache.Set("key2", 2)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key2 -> key1 -> key4" {
		t.Fatal("case 6 dlStr != \"key2 -> key1 -> key4\"")
	}

	// case 7 insert 5
	cache.Set("key5", 5)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key5 -> key2 -> key1" {
		t.Fatal("case 7 dlStr != \"key5 -> key2 -> key1\"")
	}

	// case 8 insert 1
	cache.Set("key1", 1)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key1 -> key5 -> key2" {
		t.Fatal("case 8 dlStr != \"key1 -> key5 -> key2\"")
	}

	// case 9 insert 2
	cache.Set("key2", 1)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key2 -> key1 -> key5" {
		t.Fatal("case 9 dlStr != \"key2 -> key1 -> key5\"")
	}

	// case 10 insert 3
	cache.Set("key3", 3)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key3 -> key2 -> key1" {
		t.Fatal("case 10 dlStr != \"key3 -> key2 -> key1\"")
	}

	// case 11 insert 4
	cache.Set("key4", 4)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key4 -> key3 -> key2" {
		t.Fatal("case 11 dlStr != \"key4 -> key3 -> key2\"")
	}

	// case 12 insert 5
	cache.Set("key5", 5)
	dlStr = cache.dl.print()
	fmt.Println(dlStr)
	if dlStr != "key5 -> key4 -> key3" {
		t.Fatal("case 12 dlStr != \"key5 -> key4 -> key3\"")
	}
}

func TestLruCache_Get(t *testing.T) {
	cache := NewLruCache(3)

	// case 1
	value := cache.Get("key")
	if value != nil {
		t.Fatal("TestLruCache_Get case 1 value != nil")
	}

	// case 2
	// insert 1
	cache.Set("key1", 1)
	dlStr := cache.dl.print()
	if dlStr != "key1" {
		t.Fatal("case 1 dlStr != \"key1\"")
	}
	value = cache.Get("key1")
	if value == nil {
		t.Fatal("TestLruCache_Get case 2 value == nil")
	}
	if value.(int) != 1 {
		t.Fatal("TestLruCache_Get case 2 value != 1")
	}

	// case 3 insert 2
	cache.Set("key2", 2)
	dlStr = cache.dl.print()
	if dlStr != "key2 -> key1" {
		t.Fatal("case 3 dlStr != \"key2 -> key1\"")
	}

	// case 4 insert 3
	cache.Set("key3", 3)
	dlStr = cache.dl.print()
	if dlStr != "key3 -> key2 -> key1" {
		t.Fatal("case 3 dlStr != \"key3 -> key2 -> key1\"")
	}

	// case 5 get 1
	value = cache.Get("key1")
	if value == nil {
		t.Fatal("TestLruCache_Get case 5 value == nil")
	}
	if value.(int) != 1 {
		t.Fatal("TestLruCache_Get case 5 value != 1")
	}
	dlStr = cache.dl.print()
	if dlStr != "key1 -> key3 -> key2" {
		t.Fatal("case 3 dlStr != \"key1 -> key3 -> key2\"")
	}

	// case 6 get 2
	value = cache.Get("key2")
	if value == nil {
		t.Fatal("TestLruCache_Get case 6 value == nil")
	}
	if value.(int) != 2 {
		t.Fatal("TestLruCache_Get case 6 value != 1")
	}
	dlStr = cache.dl.print()
	if dlStr != "key2 -> key1 -> key3" {
		t.Fatal("case 3 dlStr != \"key2 -> key1 -> key3\"")
	}
}
