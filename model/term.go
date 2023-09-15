package model

import (
	"database/sql"
)

type Term struct {
	SIM      string        `gorm:"primaryKey"` // 流量卡号
	Lat      float64       // 纬度
	Lng      float64       // 经度
	LocateAt sql.NullInt64 // 定位时间
	Interval sql.NullInt64 // 心跳包间隔
}
