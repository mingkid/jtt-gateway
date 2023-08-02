package admin

import (
	"net/http"

	"github.com/mingkid/jtt808-gateway/domain/service"
	"github.com/mingkid/jtt808-gateway/server/web/admin/internal/parms"
	"github.com/mingkid/jtt808-gateway/server/web/admin/internal/resp"

	"github.com/gin-gonic/gin"
)

type TermController struct {
	svr            service.Terminal
	routeGroupPath string
}

// 列表页
func (ctrl TermController) index(ctx *gin.Context) {
	dataSet, _ := ctrl.svr.All()
	var terms []resp.Term
	for _, data := range dataSet {
		terms = append(terms, resp.Term{
			SN:     data.SN,
			SIM:    data.SIM,
			Status: true,
		})
	}

	ctx.HTML(http.StatusOK, "term/index.html", gin.H{
		"title": "终端",
		"Terms": terms,
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

	// 处理表单数据
	err := ctrl.svr.Save(args.SN, args.SIM)
	if err != nil {
		ctx.String(http.StatusBadRequest, "系统异常：%s", err.Error())
	}

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
