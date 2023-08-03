package parms

// Platform 终端参数
type Platform struct {
	Identity    string `form:"identity" binding:"required"` // 平台标识
	Host        string `form:"host" binding:"required"`     // 域名
	LocationAPI string `form:"locationAPI"`                 // 定位 API
}
