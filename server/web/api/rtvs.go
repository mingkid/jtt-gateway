package api

import (
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/mingkid/jtt-gateway/log"
	"github.com/mingkid/jtt-gateway/pkg/errcode"
	"github.com/mingkid/jtt-gateway/server/jtt"
	"github.com/mingkid/jtt-gateway/server/web/api/internal/req"
	"github.com/mingkid/jtt-gateway/server/web/common"

	"github.com/gin-gonic/gin"
	engine "github.com/mingkid/g-jtt"
	"github.com/mingkid/g-jtt/protocol/bin"
	"github.com/mingkid/g-jtt/protocol/msg"
)

// videoControlAPI 视频控制API
type videoControlAPI struct {
	l log.RTVSInfoLogger
}

func NewVideoControlAPI(l log.RTVSInfoLogger) *videoControlAPI {
	return &videoControlAPI{l: l}
}

// get 请求
func (api *videoControlAPI) get(ctx *gin.Context) {
	// 默认结果
	res := RTVSResultFail
	defer func() {
		ctx.String(http.StatusOK, res)
	}()

	var params req.VideoControl
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		common.NewErrorResponse(ctx, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}

	var msgHead msg.Head
	_ = params.Bind(&msgHead)

	switch msgHead.MsgID {
	case msg.MsgIDRealtimePlay, msg.MsgIDRealtimePlayCtrl, msg.MsgIDRealtimePlayStatus:
		if err := api.send(params, msgHead); err != nil {
			if errors.Is(err, engine.DeviceOfflineError{}) {
				res = RTVSResultOffline
			}
			// TODO: 打印错误信息
			return
		}
		res = RTVSResultSuccess
	}
}

func (api *videoControlAPI) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/VideoControl", api.get)
}

func (api *videoControlAPI) send(params req.VideoControl, msgHead msg.Head) error {
	b, err := packaging(params)
	if err != nil {
		return err
	}
	if err = jtt.Svr.SendBytes(msgHead.Phone, b); err != nil {
		return err
	}
	api.l.Log(msgHead, hex.EncodeToString(b))
	return nil
}

func packaging(params req.VideoControl) ([]byte, error) {
	b, err := params.ContentBytes()
	if err != nil {
		return nil, err
	}

	// 添加校验和
	b = append(b, bin.Checksum(b))
	// 添加标识位
	b = append([]byte{126}, b...)
	b = append(b, 126)

	return b, err
}

const (
	RTVSResultFail    = "-1" // 失败
	RTVSResultOffline = "0"
	RTVSResultSuccess = "1"
)
