package conn

// ConnPool 808 设备回话缓存器
type ConnPool interface {
	// Set 设置会话
	Set(sn string, s *Connection)

	// Get 获取会话
	Get(sn string) *Connection
}
