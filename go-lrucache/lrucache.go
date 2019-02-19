// Package lrucache lru cache implement
// Created by chenguolin 2019-02-18
package lrucache

import "sync"

type hashMap map[string]interface{}

// LruCache lru cache
type LruCache struct {
	mutex    sync.Mutex //read writer mutex
	capacity int        //max cache size
	size     int        //cache used size
	dl       *list      //doubly linked list
	hMap     hashMap    //hash map
}

// NewLruCache new lru cache
func NewLruCache(maxSize int) *LruCache {
	dl := &list{}
	hMap := make(hashMap)

	return &LruCache{
		capacity: maxSize,
		size:     0,
		dl:       dl,
		hMap:     hMap,
	}
}

// Get lru cache get by key
func (lc *LruCache) Get(key string) interface{} {
	// lock
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	// check key has in hash map
	nd, ok := lc.hMap[key]
	if !ok {
		return nil
	}

	newNode := nd.(*node)

	// erase node
	lc.dl.erase(newNode)

	// insert 2 header
	lc.dl.insert(newNode)

	return newNode.value
}

// Set lru cache set key value
func (lc *LruCache) Set(key string, value interface{}) {
	// lock
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	// 1. check key in hash map
	nd, ok := lc.hMap[key]
	// key not exist in hash map
	if !ok {
		// check lrc cache full
		if lc.size == lc.capacity {
			// remove tail node
			popNd := lc.dl.pop()
			// remove hash map key
			delete(lc.hMap, popNd.key)
			// sub size
			lc.size--
		}

		// insert 2 header
		newNode := &node{
			key:   key,
			value: value,
		}
		lc.dl.insert(newNode)

		// add key 2 hash map
		lc.hMap[key] = newNode

		// add size
		lc.size++
	} else {
		// erase node
		lc.dl.erase(nd.(*node))

		// insert 2 header
		newNode := &node{
			key:   key,
			value: value,
		}
		lc.dl.insert(newNode)

		// reset hash map
		lc.hMap[key] = newNode
	}
}
