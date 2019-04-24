// Package cache data
// Modified by chenguolin 2018-03-19
package cache

import "time"

// CacheExpireTime cache expire time
const CacheExpireTime = 60

// Cache uid and address
type Cache struct {
	UIDCache     map[int]*uidCache
	AddressCache map[string]*addressCache
}

type uidCache struct {
	uid            int
	lastModifyTime int64
}

type addressCache struct {
	address        string
	lastModifyTime int64
}

var localCache *Cache

func init() {
	localCache = &Cache{
		UIDCache:     make(map[int]*uidCache),
		AddressCache: make(map[string]*addressCache),
	}
}

// AddCache add cache
func AddCache(uid int, address string) {
	nowTime := time.Now().Unix()

	uCache := &uidCache{
		uid:            uid,
		lastModifyTime: nowTime,
	}

	addrCache := &addressCache{
		address:        address,
		lastModifyTime: nowTime,
	}

	localCache.UIDCache[uid] = uCache
	localCache.AddressCache[address] = addrCache
}

// Exists check exists in cache
func Exists(uid int, address string) bool {
	nowTime := time.Now().Unix()

	if uCache, ok := localCache.UIDCache[uid]; ok {
		// cache未过期
		if uCache.lastModifyTime+CacheExpireTime> nowTime {
			return true
		}
	}

	if addrCache, ok := localCache.AddressCache[address]; ok {
		// cache未过期
		if addrCache.lastModifyTime+CacheExpireTime> nowTime {
			return true
		}
	}

	return false
}
