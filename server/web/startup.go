package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/mingkid/jtt808-gateway/server/web/api"
)

func Serve(port string) {
	// 初始化http服务
	s := gin.Default()
	fmt.Println("[web] 服务启动 :8000")
	// 注册路由
	v1.RouteRegister(s)

	_ = s.Run(port)
}
