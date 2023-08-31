package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt-gateway/pkg/errcode"
	"github.com/mingkid/jtt-gateway/server/web/api/internal/req"
	"github.com/mingkid/jtt-gateway/server/web/common"
)

// VideoControlAPI 视频控制API
type VideoControlAPI struct {
}

// get 请求
func (api VideoControlAPI) get(ctx *gin.Context) {
	params := new(req.VideoControl)
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		common.NewErrorResponse(ctx, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
	}

	msgID, err := params.MsgID()
	if err != nil {
		common.NewErrorResponse(ctx, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
	}
	switch msgID {
	case 0x9201:
		break
	case 0x9202:
		break
	case 0x920:
		break
	}
}

func (api VideoControlAPI) RegisterRoute(g *gin.RouterGroup) {
	g.GET("/VideoControl", api.get)
}
