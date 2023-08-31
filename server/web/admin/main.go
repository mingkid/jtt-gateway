package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt-gateway/domain/service"
)

var (
	termCtrl     = NewTermController(service.Terminal{})
	platformCtrl = NewPlatformController(service.Platform{})
)

func RouteRegister(g *gin.Engine) {
	termCtrl.Register(g)
	platformCtrl.Register(g)
}
