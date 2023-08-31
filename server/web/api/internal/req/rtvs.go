package req

import (
	"encoding/hex"
	"errors"

	"github.com/mingkid/g-jtt/protocol/bin"
	"github.com/mingkid/g-jtt/protocol/msg"
)

// VideoControl 视频控制参数
type VideoControl struct {
	Content                string `query:"Content"`                // 1078 协议16进制字符串(包头+包体)。包头不含7E、未转义、流水号需要808平台替换
	IsSuperiorPlatformSend bool   `query:"IsSuperiorPlatformSend"` // 是否是上级平台发送，网关可用此字段确定是否由上级平台发起。
	ChSub                  byte   `query:"ChSub"`                  // 当前通道实时子码流数量，含本次请求
	ChMain                 byte   `query:"ChMain"`                 // 当前通道实时主码流数量，含本次请求
	ChBack                 byte   `query:"ChBack"`                 // 当前通道历史流数量，含本次请求
	ChTalk                 byte   `query:"ChTalk"`                 // 当前通道对讲流数量，含本次请求
	Sub                    byte   `query:"Sub"`                    // 当前SIM下所有通道实时子码流数量，含本次请求
	Main                   byte   `query:"Main"`                   // 当前SIM下所有通道实时主码流数量，含本次请求
	Back                   byte   `query:"Back"`                   // 当前SIM下所有通道历史流数量，含本次请求
	Talk                   byte   `query:"Talk"`                   // 当前SIM下所有通道对讲流数量，含本次请求
}

// MsgID 解析 1078 消息 ID
func (req *VideoControl) MsgID() (res msg.MsgID, err error) {
	var b []byte
	if b, err = hex.DecodeString(req.Content); err != nil {
		return 0, errors.New("Content 解码错误")
	}
	return bin.ExtractMsgIDFrom(b, 0), nil
}

// ToHistoryAV 0x9201 转换消息为历史音视频传输
func (req *VideoControl) ToHistoryAV() (head *msg.Head) {
	// TODO
	return new(msg.Head)
}

// ToGetAVList 0x9205 转换消息为查询音视频列表资源
func (req *VideoControl) ToGetAVList() (head *msg.Head) {
	// TODO
	return new(msg.Head)
}

// ToHistoryVideoCtrl 0x9202 转换消息为录像回放控制
func (req *VideoControl) ToHistoryVideoCtrl() (head *msg.Head) {
	// TODO
	return new(msg.Head)
}
