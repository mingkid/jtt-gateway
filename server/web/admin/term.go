package admin

import (
	"net/http"

	"github.com/mingkid/jtt808-gateway/domain/service"
	"github.com/mingkid/jtt808-gateway/server/web/admin/internal/parms"

	"github.com/gin-gonic/gin"
)

type TermController struct {
	svr            service.Terminal
	routeGroupPath string
}

// 列表页
func (ctrl TermController) index(ctx *gin.Context) {
	terms, _ := ctrl.svr.All()
	ctx.HTML(http.StatusOK, "term/index.html", gin.H{
		"title": "Hello World",
		"terms": terms,
	})
}

// 创建页
func (ctrl TermController) create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "term/edit.html", nil)
}

// 提交
func (ctrl TermController) submit(ctx *gin.Context) {
	// 创建FormData结构体实例
	var args parms.Term

	// 将请求参数绑定到数据模型
	if err := ctx.ShouldBind(&args); err != nil {
		ctx.String(http.StatusBadRequest, "参数绑定错误：%s", err.Error())
		return
	}

	// 处理表单数据（在这里可以将数据保存到数据库，或进行其他操作）
	// 例如，打印表单数据
	println("序列号:", args.SN)
	println("SIM卡号:", args.SIM)

	// 返回响应给前端
	ctx.Redirect(http.StatusSeeOther, ctrl.routeGroupPath)
}

func NewTermController(service service.Terminal) *TermController {
	return &TermController{
		svr:            service,
		routeGroupPath: "/term",
	}
}

func (ctrl TermController) Register(g *gin.Engine) {
	group := g.Group(ctrl.routeGroupPath)
	{
		// 页面渲染 Endpoint
		group.GET("", ctrl.index)
		group.GET("/create", ctrl.create)
	}
	{
		// 接口 Endpoint
		group.POST("/submit", ctrl.submit)
	}
}
