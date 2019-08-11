package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

/* 默认缓存 - 内存 */

var (
	// DefaultMemCache 默认缓存对象
	DefaultMemCache Cache
)

func init() {
	cli := cache.New(7*24*time.Hour, 10*time.Second)
	DefaultMemCache = &MemCache{
		cli: cli,
	}
}

// MemCache 内存缓存 https://github.com/patrickmn/go-cache
type MemCache struct {
	cli *cache.Cache
}

// Get 获取一个缓存
func (mem *MemCache) Get(key string) (val string, exist bool) {
	valI, exist := mem.cli.Get(key)
	if exist == true {
		val = valI.(string)
	}
	return
}

// Set 设置一个值
func (mem *MemCache) Set(key, val string, expiration time.Duration) {
	mem.cli.Set(key, val, expiration)
}

// Del 删除知道keys
func (mem *MemCache) Del(key ...string) (err error) {
	for _, k := range key {
		mem.cli.Delete(k)
	}
	return
}
