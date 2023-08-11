package model

import "database/sql"

type Term struct {
	SN       string        `gorm:"primaryKey"` // 序列号
	SIM      string        // 流量卡号
	Lat      float64       // 纬度
	Lng      float64       // 经度
	Interval sql.NullInt32 // 心跳包间隔
}
