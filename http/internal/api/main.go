package api

import (
	"github.com/mingkid/jtt-gateway/http/internal/api/internal"
	"github.com/mingkid/jtt-gateway/log"

	"github.com/gin-gonic/gin"
)

var (
	terminal  internal.TerminalAPI
	terminals internal.TerminalsAPI
	videoCtrl *internal.VideoControlAPI
)

func RouteRegister(g *gin.Engine) {
	rootG := g.Group("/")
	{
		videoCtrl.RegisterRoute(rootG)
	}
	apiG := g.Group("/api")
	{
		terminal.Register(apiG)  // 单个终端资源注册
		terminals.Register(apiG) // 列表终端资源路由注册
	}
}

func init() {
	videoCtrl = internal.NewVideoControlAPI(log.RTVSInfoLoggerAdapter(log.RTVSInfoTermLog))
}
