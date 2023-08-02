package model

import "database/sql"

type Term struct {
	SN       string `gorm:"primaryKey"` // 序列号
	SIM      string // 流量卡号
	Interval sql.NullInt32
}
