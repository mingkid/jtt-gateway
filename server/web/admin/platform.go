package admin

import (
	"net/http"

	"github.com/mingkid/jtt808-gateway/domain/service"

	"github.com/gin-gonic/gin"
)

type PlatformController struct {
	svr            service.Platform
	routeGroupPath string
}

// 列表页面
func (ctrl PlatformController) index(ctx *gin.Context) {
	platforms, _ := ctrl.svr.All()

	// 返回响应给前端
	ctx.HTML(http.StatusOK, "platform/index.html", gin.H{
		"Title":     "业务平台",
		"platforms": platforms,
	})
}

// Register 注册控制器到指定的 Web 服务实例中
func (ctrl PlatformController) Register(g *gin.Engine) {
	group := g.Group(ctrl.routeGroupPath)
	{
		// 页面渲染 Endpoint
		group.GET("", ctrl.index)
	}
}

// NewPlatformController 初始化业务平台 Web 控制器
func NewPlatformController(svr service.Platform) PlatformController {
	return PlatformController{
		svr:            svr,
		routeGroupPath: "/platform",
	}
}
