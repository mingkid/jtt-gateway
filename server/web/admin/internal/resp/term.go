package resp

type Term struct {
	SN     string // 序号
	SIM    string // SIM 卡号
	Status bool   // 在线状态，true 为在线，false 为李现
}
