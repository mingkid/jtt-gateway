package req

// TermIdentity 终端标识选项
type TermIdentity struct {
	SIM string `uri:"sim" json:"sim" form:"sim" binding:"required"` // 序列号
}
