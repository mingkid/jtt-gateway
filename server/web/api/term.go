package api

import (
	"net/http"
	"strings"

	"github.com/mingkid/jtt808-gateway/domain/service"
	"github.com/mingkid/jtt808-gateway/server/web/api/internal/parms"
	"github.com/mingkid/jtt808-gateway/server/web/common"
	"github.com/mingkid/jtt808-gateway/server/web/common/errcode"

	"github.com/gin-gonic/gin"
)

type TerminalAPI struct{}

// post 请求；新增终端
func (api TerminalAPI) post(c *gin.Context) {
	var args parms.TermSave
	if err := c.ShouldBind(&args); err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	args.SN = strings.TrimSpace(args.SN)
	args.SIM = strings.TrimSpace(args.SIM)

	// 创建终端service
	termService := service.NewTerminal()
	err := termService.Save(args.SN, args.SIM)
	if err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	// 成功响应
	common.NewSingleResponse(c, nil).Return(http.StatusOK)
}

// delete 请求；删除终端
func (api TerminalAPI) delete(c *gin.Context) {
	var args parms.TermIdentity
	if err := c.ShouldBindUri(&args); err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	args.SN = strings.TrimSpace(args.SN)

	// 删除终端service
	termService := service.NewTerminal()
	err := termService.Delete(args.SN)
	if err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	// 成功响应
	common.NewSingleResponse(c, nil).Return(http.StatusOK)
}

func (api TerminalAPI) Register(g *gin.RouterGroup) {
	r := g.Group("/terminal")
	{
		r.POST("", api.post)         // 新增终端
		r.DELETE("/:sn", api.delete) // 删除终端
	}
}
