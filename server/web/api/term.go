package api

import (
	"net/http"
	"strings"

	"github.com/mingkid/jtt808-gateway/domain/service"
	"github.com/mingkid/jtt808-gateway/server/web/common"
	"github.com/mingkid/jtt808-gateway/server/web/common/errcode"

	"github.com/gin-gonic/gin"
)

type TerminalAPI struct{}

// Post 请求；新增终端
func (api TerminalAPI) Post(c *gin.Context) {
	var args service.TermSaveOpt
	if err := c.ShouldBindJSON(&args); err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	args.SN = strings.TrimSpace(args.SN)
	args.SIM = strings.TrimSpace(args.SIM)

	// 创建终端service
	termService := service.NewTerminal()
	err := termService.Save(args)
	if err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	// 成功响应
	common.NewSingleResponse(c, nil).Return(http.StatusOK)
}

// Delete 请求；删除终端
func (api TerminalAPI) Delete(c *gin.Context) {

}

func (api TerminalAPI) Register(g *gin.RouterGroup) {
	r := g.Group("/terminal")
	{
		r.POST("", api.Post)         // 新增终端
		r.DELETE("/:id", api.Delete) // 删除终端
	}
}
