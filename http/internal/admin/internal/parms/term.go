package parms

// Term 终端参数
type Term struct {
	SIM string `form:"sim" binding:"required"` // SIM 卡号
}
