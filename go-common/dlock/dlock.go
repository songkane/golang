// Package dlock MC实现的分布式锁
// Created by chenguolin 2018-11-16
package dlock

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	golog "gitlab.local.com/golang/go-log"
)

// LockRetryOptions 重试选项
type LockRetryOptions struct {
	Count    uint32        //重试次数
	Interval time.Duration //重试间隔
}

// MemcachedLocker MC分布式锁
type MemcacheDLocker struct {
	mcClient memcache.Client
	ctx      context.Context
}

// NewMemcacheDLocker 新建一个locker对象
func NewMemcacheDLocker(mc memcache.Client) *MemcacheDLocker {
	return &MemcacheDLocker{
		mcClient: mc,
		ctx:      context.Background(),
	}
}

// Lock 给key加锁
// key: MC对应的key
// expiration: 过期时间 单位秒
// opt: 重试选项
func (l *MemcacheDLocker) Lock(key string, expiration int32, opt LockRetryOptions) (bool, error) {
	// 默认重试次数为1
	if opt.Count <= 0 {
		opt.Count = 1
	}

	var err error
	// 重试几次
	for i := uint32(0); i < opt.Count; i++ {
		item := &memcache.Item{
			Key:        key,
			Value:      []byte{},
			Flags:      0,
			Expiration: expiration,
		}
		err = l.mcClient.Add(item)
		if err != nil {
			golog.Warn("dlock Lock mcClient Add error", golog.Err(err))
			continue
		}
		return true, nil
	}

	return false, err
}

// Unlock 释放某个key锁
// key: MC对应的key
// opt: 重试选项
func (l *MemcacheDLocker) Unlock(key string, opt LockRetryOptions) error {
	// 默认重试次数为1
	if opt.Count <= 0 {
		opt.Count = 1
	}

	var err error
	// 重试几次
	for i := uint32(0); i < opt.Count; i++ {
		err = l.mcClient.Delete(key)
		// key不存在默认返回成功
		if err == nil || err != memcache.ErrCacheMiss {
			return nil
		}
		golog.Warn("dlock Lock mcClient Add error", golog.Err(err))
		time.Sleep(opt.Interval)
		continue
	}

	return err
}
