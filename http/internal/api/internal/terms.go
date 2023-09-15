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

type TermsAPI struct {
	svr *service.Terminal
}

func NewTermsAPI() *TermsAPI {
	return &TermsAPI{
		svr: service.NewTerminal(),
	}
}

// post 请求；新增终端
func (api *TermsAPI) post(c *gin.Context) {
	var args req.TermIdentity
	if err := c.ShouldBind(&args); err != nil {
		common.NewErrorResponse(c, errcode.ParamsError.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	args.SIM = strings.TrimSpace(args.SIM)

	// 创建终端service
	err := api.svr.Save(args.SIM)
	if err != nil {
		common.NewErrorResponse(c, errcode.NotFoundError.SetMsg(err.Error())).Return(http.StatusInternalServerError)
		return
	}
	// 成功响应
	common.NewSingleResponse(c, nil).Return(http.StatusOK)
}

func (api *TermsAPI) Register(g *gin.RouterGroup) {
	r := g.Group("/terms")
	{
		r.POST("", PlatformAuth(), api.post) // 新增终端
	}
}
