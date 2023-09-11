package internal

import (
	"net/http"

	"github.com/mingkid/jtt-gateway/domain/service"
	"github.com/mingkid/jtt-gateway/http/internal/admin/internal/parms"

	"github.com/gin-gonic/gin"
)

type PlatformController struct {
	svr            service.Platform
	routeGroupPath string
}

// NewPlatformController 初始化业务平台 Web 控制器
func NewPlatformController(svr service.Platform) PlatformController {
	return PlatformController{
		svr:            svr,
		routeGroupPath: "/platform",
	}
}

// 列表页面
func (ctrl PlatformController) index(ctx *gin.Context) {
	platforms, _ := ctrl.svr.All()

	// 返回响应给前端
	ctx.HTML(http.StatusOK, "platform/index.html", gin.H{
		"Title":     "业务平台",
		"Platforms": platforms,
	})
}

// 创建页
func (ctrl PlatformController) create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "platform/edit.html", gin.H{
		"Title": "平台新增",
	})
}

// 编辑页
func (ctrl PlatformController) edit(ctx *gin.Context) {
	platform, err := ctrl.svr.GetByIdentity(ctx.Param("identity"))
	if err != nil {
		ctx.String(http.StatusNotFound, "%s", "Not Found")
	}

	// 返回响应给前端
	ctx.HTML(http.StatusOK, "platform/edit.html", gin.H{
		"Title":    "平台编辑",
		"Platform": platform,
	})
}

// 提交接口
func (ctrl PlatformController) submit(ctx *gin.Context) {
	// 创建FormData结构体实例
	var args parms.Platform

	// 将请求参数绑定到数据模型
	if err := ctx.ShouldBind(&args); err != nil {
		ctx.String(http.StatusBadRequest, "参数异常：%s", err.Error())
		return
	}

	// 处理表单数据
	err := ctrl.svr.Save(service.PlatformSaveOpt{
		Identity:    args.Identity,
		Host:        args.Host,
		LocationAPI: args.LocationAPI,
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, "系统异常：%s", err.Error())
	}

	// 返回响应给前端
	ctx.Redirect(http.StatusSeeOther, ctrl.routeGroupPath)
}

// 删除接口
func (ctrl PlatformController) del(ctx *gin.Context) {
	err := ctrl.svr.Delete(ctx.Param("ident"))
	if err != nil {
		ctx.String(http.StatusNotFound, "%s", "Not Found")
	}

	// 返回响应给前端
	ctx.Redirect(http.StatusSeeOther, ctrl.routeGroupPath)
}

// Register 注册控制器到指定的 Web 服务实例中
func (ctrl PlatformController) Register(g *gin.Engine) {
	group := g.Group(ctrl.routeGroupPath)
	{
		// 页面渲染 Endpoint
		group.GET("", ctrl.index)
		group.GET("/create", ctrl.create)
		group.GET("/edit/:identity", ctrl.edit)
	}
	{
		// 接口 Endpoint
		group.POST("/submit", ctrl.submit)
		group.GET("/del/:ident", ctrl.del)
	}
}
