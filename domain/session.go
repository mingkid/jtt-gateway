package domain

import "time"

// Session 808 设备连接会话信息
type Session struct {
	Expire time.Duration // 过期时间
}

// SessionCache 808 设备回话缓存器
type SessionCache interface {
	// Set 设置会话
	Set(sn string, s *Session)

	// Get 获取会话
	Get(sn string) *Session
}

// MemorySessionCache 808 设备连接会话内存缓存器
type MemorySessionCache struct {
	cache map[string]*Session
}

func (m *MemorySessionCache) Set(sn string, s *Session) {
	m.cache[sn] = s
}

func (m *MemorySessionCache) Get(sn string) *Session {
	return m.cache[sn]
}

func NewMemorySessionCache() *MemorySessionCache {
	return &MemorySessionCache{
		cache: make(map[string]*Session, 100),
	}
}
