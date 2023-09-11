package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt-gateway/domain/service"
	"github.com/mingkid/jtt-gateway/http/internal/admin/internal"
)

var (
	termCtrl     = internal.NewTermController(service.Terminal{})
	platformCtrl = internal.NewPlatformController(service.Platform{})
)

func RouteRegister(g *gin.Engine) {
	termCtrl.Register(g)
	platformCtrl.Register(g)
}
