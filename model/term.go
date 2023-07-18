package model

import "database/sql"

type Term struct {
	SN    string // 序列号
	SIM   string // 流量卡号
	Alive sql.NullBool
}
