package parms

// Term 终端参数
type Term struct {
	SN  string `form:"sn" binding:"required"` // 序号
	SIM string `form:"sn"`                    // SIM 卡号
}
