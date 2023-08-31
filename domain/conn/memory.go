package conn

import (
	"sync"
)

// MemoryConnPool JTT 设备连接池
type MemoryConnPool struct {
	conns map[string]*Connection
	mu    sync.Mutex
}

func (m *MemoryConnPool) Set(sn string, conn *Connection) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.conns[sn] = conn
}

func (m *MemoryConnPool) Get(sn string) *Connection {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.disconnectOnTimeout(sn)
	return m.conns[sn]
}

func (m *MemoryConnPool) disconnectOnTimeout(sn string) {
	if c, found := m.conns[sn]; found {
		if c.IsTimeout() {
			delete(m.conns, sn)
		}
	}
}

var (
	defaultConnPool     *MemoryConnPool
	defaultConnPoolOnce sync.Once
)

func DefaultConnPool() *MemoryConnPool {
	defaultConnPoolOnce.Do(func() {
		if defaultConnPool == nil {
			defaultConnPool = &MemoryConnPool{
				conns: make(map[string]*Connection),
			}
		}
	})
	return defaultConnPool
}
