package req

import (
	"encoding/hex"

	"github.com/mingkid/g-jtt/protocol/bin"
	"github.com/mingkid/g-jtt/protocol/codec"
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

// Bind 绑定传递的结构体指针
func (vc *VideoControl) Bind(m interface{}) error {
	b, err := vc.ContentBytes()
	if err != nil {
		return err
	}

	b = bin.Unescape(b)

	d := new(codec.Decoder)
	return d.Decode(m, b)
}

func (vc *VideoControl) ContentBytes() ([]byte, error) {
	return hex.DecodeString(vc.Content)
}
