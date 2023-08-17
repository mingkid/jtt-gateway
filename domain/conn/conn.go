package conn

import (
	"net"
	"time"
)

// Connection 808 设备连接会话信息
type Connection struct {
	expire time.Time // 过期时间戳
	conn   net.Conn  // 连接对象
}

func NewConnection(conn net.Conn, expire time.Time) *Connection {
	return &Connection{
		expire: expire,
		conn:   conn,
	}
}

// IsTimeout 超时
func (c *Connection) IsTimeout() bool {
	return time.Now().Unix() <= c.expire.Unix()
}

// SetExpireDurationFromNow 设置距离现在多久后过期
func (c *Connection) SetExpireDurationFromNow(d time.Duration) {
	c.SetExpireForTimestamp(int64(d) + time.Now().Unix())
}

// SetExpireForTimestamp 通过时间戳设置过期时间
func (c *Connection) SetExpireForTimestamp(s int64) {
	c.SetExpire(time.Unix(s, 0))
}

// SetExpire 设置过期时间
func (c *Connection) SetExpire(t time.Time) {
	c.expire = t
}

// Send 发送数据
func (c *Connection) Send(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}

// Receipt 接收数据
func (c *Connection) Receipt() ([]byte, error) {
	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	return b[:n], err
}
