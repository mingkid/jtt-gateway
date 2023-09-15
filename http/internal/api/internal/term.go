package internal

import (
	"net/http"
	"strings"

	"github.com/mingkid/jtt-gateway/domain/service"
	"github.com/mingkid/jtt-gateway/http/internal/api/internal/req"
	"github.com/mingkid/jtt-gateway/http/internal/common"
	"github.com/mingkid/jtt-gateway/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type TermAPI struct {
	svr *service.Terminal
}

func NewTermAPI() *TermAPI {
	return &TermAPI{
		svr: service.NewTerminal(),
	}
}

// delete 请求；删除终端
func (api *TermAPI) delete(c *gin.Context) {
	var args req.TermIdentity
	if err := c.ShouldBindUri(&args); err != nil {
		common.NewErrorResponse(c, errcode.ParamsError.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	args.SIM = strings.TrimSpace(args.SIM)

	// 删除终端service
	err := api.svr.Delete(args.SIM)
	if err != nil {
		common.NewErrorResponse(c, errcode.NotFoundError.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	// 成功响应
	common.NewSingleResponse(c, nil).Return(http.StatusOK)
}

func (api *TermAPI) Register(g *gin.RouterGroup) {
	r := g.Group("/term")
	{
		r.DELETE("/:sn", PlatformAuth(), api.delete) // 删除终端
	}
}
