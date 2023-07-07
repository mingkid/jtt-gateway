package log

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

type InfoFormatter func(params InfoFormatterParams) string

type InfoFormatterParams struct {
	Time   time.Time
	IP     net.Addr
	Result uint8
	Phone  string
	MsgID  []byte
	Data   []byte
}

func (p InfoFormatterParams) ResultColor() string {
	if p.Result > 0 {
		return "\033[90;43m"
	}
	return "\033[97;42m"
}

func (p InfoFormatterParams) ResetColor() string {
	return "\033[0m"
}

func DefaultInfoFormatter(params InfoFormatterParams) string {
	result := "Success"
	if params.Result > 0 {
		result = "Failed"
	}
	return fmt.Sprintf("[JTT808] %v | %s | %15s ｜%s %v %s｜%s | %s \n",
		params.Time.Format("2006/01/02 - 15:04:05"),
		params.IP,
		params.Phone,
		params.ResultColor(), result, params.ResetColor(),
		hex.EncodeToString(params.MsgID),
		hex.EncodeToString(params.Data),
	)
}
