package cache

import (
	"fmt"
	"time"
)

/* 缓存对象 */

// Cache 缓存接口 - 用户缓存登录信息
type Cache interface {
	// Get 获取一个缓存
	Get(key string) (val string, exist bool)
	// Set 设置一个值
	Set(key, val string, expiration time.Duration)
	// Del 删除知道keys
	Del(key ...string) (err error)
}

// 常量
const (
	LoginKey = "login:%s"
)

// GetLoginKey 获取登录key
func GetLoginKey(token string) string {
	return fmt.Sprintf(LoginKey, token)
}
