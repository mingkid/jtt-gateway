package http

import (
	"github.com/mingkid/jtt-gateway/http/internal/admin"
	v1 "github.com/mingkid/jtt-gateway/http/internal/api"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	// 初始化http服务
	s := gin.Default()
	s.LoadHTMLGlob("template/**/*")

	// 注册路由
	v1.RouteRegister(s)
	admin.RouteRegister(s)

	_ = s.Run(port)
}
