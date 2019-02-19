// Package lrucache unit test
// Created by chenguolin 2019-02-18
package lrucache

import (
	"fmt"
	"testing"
)

func TestNode_insert(t *testing.T) {
	dl := &list{}

	// case 1
	listStr := dl.print()
	if listStr != "" {
		t.Fatal("TestNode_insert case 1 listStr != \"\"")
	}

	// case 2 insert node 1
	n1 := &node{
		key:   "node1-key",
		value: "node1-value",
	}
	dl.insert(n1)
	listStr = dl.print()
	if listStr != "node1-key" {
		t.Fatal("TestNode_insert case 2 failed ~")
	}
	fmt.Println(dl.print())

	// case 3 insert node 2
	n2 := &node{
		key:   "node2-key",
		value: "node2-value",
	}
	dl.insert(n2)
	listStr = dl.print()
	if listStr != "node2-key -> node1-key" {
		t.Fatal("TestNode_insert case 3 failed ~")
	}
	fmt.Println(dl.print())

	// case 4 insert node 3
	n3 := &node{
		key:   "node3-key",
		value: "node3-value",
	}
	dl.insert(n3)
	listStr = dl.print()
	if listStr != "node3-key -> node2-key -> node1-key" {
		t.Fatal("TestNode_insert case 4 failed ~")
	}
	fmt.Println(dl.print())

	// case 5 insert node 4
	n4 := &node{
		key:   "node4-key",
		value: "node4-value",
	}
	dl.insert(n4)
	listStr = dl.print()
	if listStr != "node4-key -> node3-key -> node2-key -> node1-key" {
		t.Fatal("TestNode_insert case 5 failed ~")
	}
	fmt.Println(dl.print())
}

func TestNode_erase(t *testing.T) {
	dl := &list{}

	// node1
	n1 := &node{
		key:   "node1-key",
		value: "node1-value",
	}
	dl.insert(n1)
	// node 2
	n2 := &node{
		key:   "node2-key",
		value: "node2-value",
	}
	dl.insert(n2)
	// node 3
	n3 := &node{
		key:   "node3-key",
		value: "node3-value",
	}
	dl.insert(n3)
	// node 4
	n4 := &node{
		key:   "node4-key",
		value: "node4-value",
	}
	dl.insert(n4)

	// case 1 erase node 2
	dl.erase(n2)
	listStr := dl.print()
	if listStr != "node4-key -> node3-key -> node1-key" {
		t.Fatal("TestNode_erase case 1 failed ~")
	}
	fmt.Println(dl.print())

	// case 2 erase node 4
	dl.erase(n4)
	listStr = dl.print()
	if listStr != "node3-key -> node1-key" {
		t.Fatal("TestNode_erase case 2 failed ~")
	}
	fmt.Println(dl.print())

	// case 3 erase node 1
	dl.erase(n1)
	listStr = dl.print()
	if listStr != "node3-key" {
		t.Fatal("TestNode_erase case 3 failed ~")
	}
	fmt.Println(dl.print())

	// case 4 erase node 3
	dl.erase(n3)
	listStr = dl.print()
	if listStr != "" {
		t.Fatal("TestNode_erase case 4 failed ~")
	}
	fmt.Println(dl.print())
}

func TestNode_pop(t *testing.T) {
	dl := &list{}

	// node1
	n1 := &node{
		key:   "node1-key",
		value: "node1-value",
	}
	dl.insert(n1)
	// node 2
	n2 := &node{
		key:   "node2-key",
		value: "node2-value",
	}
	dl.insert(n2)
	// node 3
	n3 := &node{
		key:   "node3-key",
		value: "node3-value",
	}
	dl.insert(n3)
	// node 4
	n4 := &node{
		key:   "node4-key",
		value: "node4-value",
	}
	dl.insert(n4)

	// case 1
	nd := dl.pop()
	if nd == nil {
		t.Fatal("TestNode_erase case 1 nd == nil failed ~")
	}
	listStr := dl.print()
	if listStr != "node4-key -> node3-key -> node2-key" {
		t.Fatal("TestNode_erase case 1 failed ~")
	}
	fmt.Println(dl.print())

	// case 2
	nd = dl.pop()
	if nd == nil {
		t.Fatal("TestNode_erase case 2 nd == nil failed ~")
	}
	listStr = dl.print()
	if listStr != "node4-key -> node3-key" {
		t.Fatal("TestNode_erase case 2 failed ~")
	}
	fmt.Println(dl.print())

	// case 3
	nd = dl.pop()
	if nd == nil {
		t.Fatal("TestNode_erase case 3 nd == nil failed ~")
	}
	listStr = dl.print()
	if listStr != "node4-key" {
		t.Fatal("TestNode_erase case 3 failed ~")
	}
	fmt.Println(dl.print())

	// case 4
	nd = dl.pop()
	if nd == nil {
		t.Fatal("TestNode_erase case 4 nd == nil failed ~")
	}
	listStr = dl.print()
	if listStr != "" {
		t.Fatal("TestNode_erase case 4 failed ~")
	}
	fmt.Println(dl.print())
}
