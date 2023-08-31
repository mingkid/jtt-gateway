package publish

import (
	"bytes"
	"encoding/json"

	"github.com/mingkid/g-jtt/protocol/msg"
)

type LocationOpt struct {
	Phone      string `json:"phoneNumber"` // 手机号
	Longitude  uint32 `json:"longitude"`   // 经度
	Latitude   uint32 `json:"latitude"`    // 纬度
	Altitude   uint16 `json:"altitude"`    // 高程
	Speed      uint16 `json:"speed"`       // 速度
	Direction  uint16 `json:"direction"`   // 方向
	AccOn      bool   `json:"accOn"`       // ACC状态，true 表示开，false 表示关
	LocateAt   int64  `json:"locateAt"`    // 定位时间戳
	IsReReport bool   `json:"isReReport"`  // 是否为补传数据

	Warning    `json:"warning"`    // 报警标志
	Status     `json:"status"`     // 状态
	M0200Extra `json:"additional"` // 附加信息
}

func (l LocationOpt) Buffer() (*bytes.Buffer, error) {
	locationJson, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(locationJson), nil
}

func NewLocationOpt(phone string, msg *msg.M0200, isReReport bool) LocationOpt {
	t, _ := msg.LocateTime()
	return LocationOpt{
		Phone:      phone,
		Longitude:  msg.Longitude,
		Latitude:   msg.Latitude,
		Altitude:   msg.Altitude,
		Speed:      msg.Speed,
		Direction:  msg.Direction,
		LocateAt:   t.Unix(),
		IsReReport: isReReport,
		Warning:    NewWarning(msg.Warn),
		Status:     NewStatus(msg.Status),
		M0200Extra: NewExtra(msg.Extra),
	}
}

// Warning 报警标志
type Warning struct {
	Emergency              bool `json:"emergency"`              // 紧急报警
	OverSpeed              bool `json:"overSpeed"`              // 超速报警
	Tired                  bool `json:"tired"`                  // 疲劳驾驶报警
	Danger                 bool `json:"danger"`                 // 危险预警
	TermFault              bool `json:"termFault"`              // 终端故障
	AerialUnConn           bool `json:"aerialUnConn"`           // GSM模块故障
	AerialShortCircuit     bool `json:"aerialShortCircuit"`     // GPS模块故障
	TermUndervoltage       bool `json:"termUndervoltage"`       // 终端主电源欠压
	TermPowerFail          bool `json:"termPowerFail"`          // 终端主电源掉电
	LCDFault               bool `json:"lcdFault"`               // 终端LCD或显示器故障
	TTSFault               bool `json:"ttsFault"`               // TTS模块故障
	CameraFault            bool `json:"cameraFault"`            // 摄像头故障
	ICCardFault            bool `json:"icCardFault"`            // IC卡模块故障
	PreOverSpeed           bool `json:"preOverSpeed"`           // 预警超速
	PreTired               bool `json:"preTired"`               // 预警疲劳驾驶
	TotalDriveTimeout      bool `json:"totalDriveTimeout"`      // 累计驾驶时长超时
	StopTimeout            bool `json:"stopTimeout"`            // 停车超时
	AreaIO                 bool `json:"areaIO"`                 // 区域进出报警
	RoadIO                 bool `json:"roadIO"`                 // 路线进出报警
	RoadDriveNotEnoughTime bool `json:"roadDriveNotEnoughTime"` // 路线驾驶时长不足
	RoadDeparture          bool `json:"roadDeparture"`          // 路线偏离报警
	VSSFault               bool `json:"vssFault"`               // 车辆VSS故障
	OilError               bool `json:"oilError"`               // 油量异常报警
	VehStolen              bool `json:"vehStolen"`              // 车辆被盗报警
	IllegalIgnition        bool `json:"illegalIgnition"`        // 非法点火报警
	IllegalMove            bool `json:"illegalMove"`            // 非法位移报警
	PreCrash               bool `json:"preCrash"`               // 碰撞预警
	PreRollOver            bool `json:"preRollOver"`            // 侧翻预警
	IllegalOpenDoor        bool `json:"illegalOpenDoor"`        // 非法开门报警
}

func NewWarning(warn msg.M0200Warn) Warning {
	return Warning{
		Emergency:              warn.Emergency(),
		OverSpeed:              warn.OverSpeed(),
		Tired:                  warn.Tired(),
		Danger:                 warn.Danger(),
		TermFault:              warn.TermFault(),
		AerialUnConn:           warn.AerialUnConn(),
		AerialShortCircuit:     warn.AerialShortCircuit(),
		TermUndervoltage:       warn.TermUndervoltage(),
		TermPowerFail:          warn.TermPowerFail(),
		LCDFault:               warn.LCDFault(),
		TTSFault:               warn.TTSFault(),
		CameraFault:            warn.CameraFault(),
		ICCardFault:            warn.ICCardFault(),
		PreOverSpeed:           warn.PreOverSpeed(),
		PreTired:               warn.PreTired(),
		TotalDriveTimeout:      warn.TotalDriveTimeout(),
		StopTimeout:            warn.StopTimeout(),
		AreaIO:                 warn.AreaIO(),
		RoadIO:                 warn.RoadIO(),
		RoadDriveNotEnoughTime: warn.RoadDriveNotEnoughTime(),
		RoadDeparture:          warn.RoadDeparture(),
		VSSFault:               warn.VSSFault(),
		OilError:               warn.OilError(),
		VehStolen:              warn.VehStolen(),
		IllegalIgnition:        warn.IllegalIgnition(),
		IllegalMove:            warn.IllegalMove(),
		PreCrash:               warn.PreCrash(),
		PreRollOver:            warn.PreRollOver(),
		IllegalOpenDoor:        warn.IllegalOpenDoor(),
	}
}

// Status 状态位
type Status struct {
	ACCOn               bool  `json:"accOn"`               // 0: ACC 开关
	PositionOn          bool  `json:"positionOn"`          // 1: 定位开
	IsSLat              bool  `json:"isSLat"`              // 2: 0：北纬；1：南纬
	IsWLong             bool  `json:"isWLong"`             // 3: 0：东经；1：西经
	InOperation         bool  `json:"inOperation"`         // 4: 0：运营状态；1：停运状态
	Encrypted           bool  `json:"encrypted"`           // 5: 0：经纬度未经保密插件加密；1：经纬度已经保密插件加密
	LoadStatus          uint8 `json:"loadStatus"`          // 8-9: 00：空车；01：半载；10：保留；11：满载
	IsOilChannelNormal  bool  `json:"isOilChannelNormal"`  // 10: 0：车辆油路正常；1：车辆油路断开
	IsCircuitNormal     bool  `json:"isCircuitNormal"`     // 11: 0：车辆电路正常；1：车辆电路断开
	DoorLocked          bool  `json:"doorLocked"`          // 12: 0：车门解锁；1：车门加锁
	FrontDoorOpened     bool  `json:"frontDoorOpened"`     // 13: 0：门1关；1：门1开（前门）
	MidDoorOpened       bool  `json:"midDoorOpened"`       // 14: 0：门2关；1：门2开（中门）
	BackDoorOpened      bool  `json:"backDoorOpened"`      // 15: 0：门3关；1：门3开（后门）
	DriveRoomDoorOpened bool  `json:"driveRoomDoorOpened"` // 16: 0：门4关；1：门4开（驾驶席门）
	ElseRoomDoorOpened  bool  `json:"elseRoomDoorOpened"`  // 17: 0：门5关；1：门5开（自定义）
	GPSUsed             bool  `json:"gpsUsed"`             // 18: 0：未使用GPS卫星进行定位；1：使用GPS卫星进行定位
	BeiDouUsed          bool  `json:"beiDouUsed"`          // 19: 0：未使用北斗卫星进行定位；1：使用北斗卫星进行定位
	GLONASSUsed         bool  `json:"gLONASSUsed"`         // 20: 0：未使用GLONASS卫星进行定位；1：使用GLONASS卫星进行定位
	GalileoUsed         bool  `json:"galileoUsed"`         // 21: 0：未使用Galileo卫星进行定位；1：使用Galileo卫星进行定位
}

func NewStatus(status msg.M0200Status) Status {
	return Status{
		ACCOn:               status.ACCOn(),
		PositionOn:          status.PositionOn(),
		IsSLat:              status.IsSLat(),
		IsWLong:             status.IsWLong(),
		InOperation:         status.Operating(),
		Encrypted:           status.Encrypted(),
		LoadStatus:          0,
		IsOilChannelNormal:  status.OilChannelNormal(),
		IsCircuitNormal:     status.CircuitNormal(),
		DoorLocked:          status.DoorLocked(),
		FrontDoorOpened:     status.FrontDoorOpened(),
		MidDoorOpened:       status.MidDoorOpened(),
		BackDoorOpened:      status.BackDoorOpened(),
		DriveRoomDoorOpened: status.DriveRoomDoorOpened(),
		ElseRoomDoorOpened:  status.ElseRoomDoorOpened(),
		GPSUsed:             status.GPSUsed(),
		BeiDouUsed:          status.BPSUsed(),
		GLONASSUsed:         status.GLONASSUsed(),
		GalileoUsed:         status.GalileoUsed(),
	}
}

// M0200Extra 附加信息
type M0200Extra struct {
	Mileage        uint32 `json:"mileage"`        // 里程（单位：公里）
	Oil            uint16 `json:"oil"`            // 油量（单位：升）
	Speed          uint16 `json:"speed"`          // 速度（单位：km/h）
	ConformWarnID  uint16 `json:"conformWarnID"`  // 预警标识号
	Analog         uint32 `json:"analog"`         // 模拟量
	SignalStrength byte   `json:"signalStrength"` // 信号强度
	GNSSQty        byte   `json:"gnssQty"`        // GNSS数量

	OverSpeedWarn     `json:"overSpeedWarn"`     // 超速报警
	IOPositionWarn    `json:"ioPositionWarn"`    // 进出区域/路线报警
	DriveTooShortWarn `json:"driveTooShortWarn"` // 路线行驶时间不足/过长报警附加信息
	VehSignalStatus   `json:"vehSignalStatus"`   // 扩展车辆信号状态位 (ID: 0x25, 长度: 4 字节)。包含扩展车辆信号状态位的定义
	IOStatus          `json:"ioStatus"`          // IO 状态位 (ID: 0x2A, 长度: 2 字节)。包含 IO 状态位的定义
}

func NewExtra(extra msg.M0200Extra) M0200Extra {
	return M0200Extra{
		Mileage:           extra.Mileage(),
		Oil:               extra.Oil(),
		Speed:             extra.Speed(),
		ConformWarnID:     extra.ConformWarnID(),
		Analog:            extra.Analog(),
		SignalStrength:    extra.SignalStrength(),
		GNSSQty:           extra.GNSSQty(),
		OverSpeedWarn:     NewOverSpeedWarn(extra.OverSpeedWarn()),
		IOPositionWarn:    NewIOPositionWarn(extra.IOPositionWarn()),
		DriveTooShortWarn: NewDriveTooShortWarn(extra.DriveTooShortWarn()),
		VehSignalStatus:   NewVehSignalStatus(extra.VehSignalStatus()),
		IOStatus:          NewIOStatus(extra.IOStatus()),
	}
}

// OverSpeedWarn 超速报警附加信息
type OverSpeedWarn struct {
	PositionType byte   `json:"positionType"` // 位置类型
	PositionID   uint32 `json:"positionID"`   // 区域或路段ID
}

func NewOverSpeedWarn(warn msg.OverSpeedWarn) OverSpeedWarn {
	return OverSpeedWarn{
		PositionType: byte(warn.PositionType()),
		PositionID:   warn.PositionID(),
	}
}

// IOPositionWarn 进出区域/路线报警附加信息
type IOPositionWarn struct {
	PositionType byte   `json:"positionType"` // 位置类型
	PositionID   uint32 `json:"positionID"`   // 区域或线路ID
	Direction    byte   `json:"direction"`    // 方向
}

func NewIOPositionWarn(warn msg.IOPositionWarn) IOPositionWarn {
	return IOPositionWarn{
		PositionType: byte(warn.PositionType()),
		PositionID:   warn.PositionID(),
		Direction:    byte(warn.Direction()),
	}
}

// DriveTooShortWarn 路线行驶时间不足/过长报警附加信息
type DriveTooShortWarn struct {
	RoadID    uint32 `json:"roadID"`    // 路段ID
	DriveTime uint16 `json:"driveTime"` // 路段行驶时间
	Result    byte   `json:"result"`    // 结果
}

func NewDriveTooShortWarn(warn msg.DriveTooShortWarn) DriveTooShortWarn {
	return DriveTooShortWarn{
		RoadID:    warn.RoadID(),
		DriveTime: warn.DriveTime(),
		Result:    warn.Result(),
	}
}

// VehSignalStatus 扩展车辆信号状态位
type VehSignalStatus struct {
	LowBeam         bool `json:"lowBeam"`         // 近光灯
	HighBeam        bool `json:"highBeam"`        // 远光灯
	RightTurnSignal bool `json:"rightTurnSignal"` // 右转灯
	LeftTurnSignal  bool `json:"leftTurnSignal"`  // 左转灯
	Break           bool `json:"break"`           // 刹车
	Reverse         bool `json:"reverse"`         // 倒档
	FogLights       bool `json:"fogLights"`       // 雾灯
	PositionLights  bool `json:"positionLights"`  // 示廓灯
	Horn            bool `json:"horn"`            // 喇叭
	AirConditioning bool `json:"airConditioning"` // 空调状态
	Neutral         bool `json:"neutral"`         // 空档
	Retarder        bool `json:"retarder"`        // 缓速器
	ABS             bool `json:"abs"`             // 防抱死制动系统
	Heater          bool `json:"heater"`          // 加热器
	Clutch          bool `json:"clutch"`          // 离合器
}

func NewVehSignalStatus(status msg.VehSignalStatus) VehSignalStatus {
	return VehSignalStatus{
		LowBeam:         status.LowBeam(),
		HighBeam:        status.HighBeam(),
		RightTurnSignal: status.RightTurnSignal(),
		LeftTurnSignal:  status.LeftTurnSignal(),
		Break:           status.Break(),
		Reverse:         status.Reverse(),
		FogLights:       status.FogLights(),
		PositionLights:  status.PositionLights(),
		Horn:            status.Horn(),
		AirConditioning: status.AirConditioning(),
		Neutral:         status.Neutral(),
		Retarder:        status.Retarder(),
		ABS:             status.ABS(),
		Heater:          status.Heater(),
		Clutch:          status.Clutch(),
	}
}

// IOStatus IO状态位
type IOStatus struct {
	DeepSleep bool `json:"deepSleep"` // 深度休眠
	Sleep     bool `json:"sleep"`     // 休眠
}

func NewIOStatus(status msg.IOStatus) IOStatus {
	return IOStatus{
		DeepSleep: status.DeepSleep(),
		Sleep:     status.Sleep(),
	}
}
