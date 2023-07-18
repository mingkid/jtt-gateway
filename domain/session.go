package domain

import (
	"time"
)

// TODO
// 1、创建session类
// 2、使用init来引用初始化创建他
// 3、每次发送心跳包就更新对应的缓存
// 4、读不到数据就默认离线

var HeartBeat *Session

type Session struct {
	cache map[string]any
}

func (t *Session) Get(key string) bool {
	value := t.cache[key]

	if value == nil {
		return false
	}
	e, ok := value.(time.Time)
	if !ok {
		return false
	}

	if time.Now().Before(e) {
		// 过期
		delete(t.cache, key)
		return false
	}
	return true
}

func (t *Session) Set(key string, expire time.Duration) {
	t.cache[key] = expire
}

func NewSession() *Session {
	return &Session{
		cache: map[string]any{},
	}
}

func init() {
	HeartBeat = NewSession()
}
