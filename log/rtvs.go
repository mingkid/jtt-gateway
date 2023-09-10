package log

import (
	"fmt"
	"time"

	"github.com/mingkid/g-jtt/protocol/msg"
)

// RTVSInfoLogger RTVS 日志接口
type RTVSInfoLogger interface {
	Log(h msg.Head, content string)
}

// RTVSInfoLoggerAdapter 适配器
type RTVSInfoLoggerAdapter func(h msg.Head, content string)

// Log 日志记录
func (l RTVSInfoLoggerAdapter) Log(h msg.Head, content string) {
	l(h, content)
}

// RTVSInfoTermLog 终端日志
func RTVSInfoTermLog(h msg.Head, content string) {
	format := "[JTT-Gateway-RTVS] %s | %s --> %s\n%s\n"
	opts := []any{
		time.Now().Format("2006/01/02 - 15:04:05"),
		h.MsgID.Hex(),
		h.Phone,
		content,
	}
	fmt.Printf(format, opts...)
}
