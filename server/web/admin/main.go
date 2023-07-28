package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt808-gateway/domain/service"
)

var (
	index = NewTermController(service.Terminal{})
)

func RouteRegister(g *gin.Engine) {
	index.Register(g)
}
