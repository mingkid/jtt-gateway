package model

type Platform struct {
	ID          uint
	Identity    string // 平台标识
	Host        string // 域名
	LocationAPI string // 定位API
	AccessToken string // 访问 Token
}
