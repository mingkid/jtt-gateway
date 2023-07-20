package model

type Platform struct {
	ID          uint
	Identity    string // 平台名称
	Host        string // 域名
	LocationAPI string // 定位API
	Method      string // 请求方式
}
