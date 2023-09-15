package api

import (
	"github.com/mingkid/jtt-gateway/http/internal/api/internal"
	"github.com/mingkid/jtt-gateway/log"

	"github.com/gin-gonic/gin"
)

var (
	termAPI   *internal.TermAPI
	termsAPI  *internal.TermsAPI
	videoCtrl *internal.VideoControlAPI
)

func RouteRegister(g *gin.Engine) {
	api := g.Group("/api")
	{
		videoCtrl.RegisterRoute(api) // RTVS 资源路由注册
		termAPI.Register(api)        // 单个终端资源注册
		termsAPI.Register(api)       // 列表终端资源路由注册
	}
}

func init() {
	termAPI = internal.NewTermAPI()
	termsAPI = internal.NewTermsAPI()
	videoCtrl = internal.NewVideoControlAPI(log.RTVSInfoLoggerAdapter(log.RTVSInfoTermLog))
}
