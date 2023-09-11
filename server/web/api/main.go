package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt-gateway/log"
)

var (
	terminal  TerminalAPI
	terminals TerminalsAPI
	videoCtrl *videoControlAPI
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
	videoCtrl = NewVideoControlAPI(log.RTVSInfoLoggerAdapter(log.RTVSInfoTermLog))
}
