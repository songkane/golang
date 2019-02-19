// Package lrucache double doubly linked list implement
// Created by chenguolin 2019-02-18
package lrucache

// list doubly linked list
type list struct {
	header *node //list header node
	size   int   //list size
}

// node list node
type node struct {
	pre   *node
	next  *node
	key   string
	value interface{}
}

// insert node 2 header
func (l *list) insert(n *node) {
	// check is empty list
	if l.header == nil {
		n.pre = n
		n.next = n
		l.header = n
		l.size++
		return
	}

	tmpHeader := l.header
	tmpTail := l.header.pre

	// set n pre and next pointer
	n.next = tmpHeader
	n.pre = tmpTail

	// set tmpHeader pre
	tmpHeader.pre = n

	// set tmpTail next
	tmpTail.next = n

	// reset header
	l.header = n
	l.size++
	return
}

// erase node
func (l *list) erase(n *node) {
	if l.header == nil {
		return
	}

	// iterator list
	tmpNode := l.header
	iter := 1

	for {
		if iter > l.size {
			break
		}

		// check found
		if tmpNode.key == n.key {
			if l.size == 1 {
				l.header = nil
				l.size--
			} else {
				// reset tmpNode pre
				tmpNode.pre.next = tmpNode.next
				// reset tmpNode next
				tmpNode.next.pre = tmpNode.pre
				// sub size
				l.size--

				if iter == 1 {
					l.header = tmpNode.next
				}
			}

			break
		}

		// next
		tmpNode = tmpNode.next
		iter++
	}
}

// pop tail node
func (l *list) pop() *node {
	if l.header == nil {
		return nil
	}

	var popNode *node

	if l.size == 1 {
		popNode = l.header

		l.header = nil
		l.size--
	} else if l.size == 2 {
		popNode = l.header.next

		l.header.next = l.header
		l.header.pre = l.header
		l.size--
	} else {
		popNode = l.header.pre

		tmpHeader := l.header
		tailNode := l.header.pre
		// reset header pre
		tmpHeader.pre = tailNode.pre
		// reset tail node pre node's next
		tailNode.pre.next = tmpHeader
		// sub size
		l.size--
	}

	return popNode
}

// print list all node
func (l *list) print() string {
	if l.header == nil {
		return ""
	}
	// iterator list
	tmpNode := l.header
	iter := 1
	listStr := ""

	for {
		if iter > l.size {
			break
		}

		listStr += tmpNode.key
		if tmpNode.next != l.header {
			listStr += " -> "
		}

		tmpNode = tmpNode.next
		iter++
	}

	return listStr
}
