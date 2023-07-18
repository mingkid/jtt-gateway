package api

import (
	"github.com/gin-gonic/gin"
)

var (
	terminal  = new(TerminalAPI)
	terminals = new(TerminalsAPI)
)

func RouteRegister(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		terminal.Register(r)  // 单个终端资源注册
		terminals.Register(r) // 列表终端资源路由注册
	}
}
