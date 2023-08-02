package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt808-gateway/domain/service"
)

var (
	termIndex = NewTermController(service.Terminal{})
)

func RouteRegister(g *gin.Engine) {
	termIndex.Register(g)
}
