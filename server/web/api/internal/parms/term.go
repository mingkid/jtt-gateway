package parms

// TermIdentity 终端标识选项
type TermIdentity struct {
	SN string `json:"sn" binding:"required"` // 序列号
}

// TermSave 终端保存选项
type TermSave struct {
	TermIdentity
	SIM string `json:"sim"` // 流量卡号
}
