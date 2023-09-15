package resp

type Term struct {
	SIM      string  // SIM 卡号
	Status   bool    // 在线状态，true 为在线，false 为离线
	Lng      float64 // 经度
	Lat      float64 // 纬度
	LocateAt int64   // 定位时间
}
