package http

import (
	"github.com/mingkid/jtt-gateway/http/internal/admin"
	v1 "github.com/mingkid/jtt-gateway/http/internal/api"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	// 初始化http服务
	s := gin.Default()

	// 注册路由
	v1.RouteRegister(s)
	admin.RouteRegister(s)

	// 加载模板
	s.LoadHTMLGlob("template/**/*")

	_ = s.Run(port)
}
