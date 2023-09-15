package parms

type PlatformIdentity struct {
	Identity string `uri:"ident" form:"ident" binding:"required"` // 平台标识
}

// Platform 终端参数
type Platform struct {
	PlatformIdentity
	Host        string `form:"host" binding:"required"` // 域名
	LocationAPI string `form:"locationAPI"`             // 定位 API
}
