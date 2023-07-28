package admin

import (
	"net/http"

	"github.com/mingkid/jtt808-gateway/domain/service"

	"github.com/gin-gonic/gin"
)

type TermController struct {
	svr service.Terminal
}

func (ctrl TermController) index(ctx *gin.Context) {
	terms, _ := ctrl.svr.All()
	ctx.HTML(http.StatusOK, "term/index.html", gin.H{
		"title": "Hello World",
		"terms": terms,
	})
}

func NewTermController(service service.Terminal) *TermController {
	return &TermController{
		svr: service,
	}
}

func (ctrl TermController) Register(g *gin.Engine) {
	g.GET("", ctrl.index)
}
