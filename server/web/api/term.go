package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt808-gateway/domain/service"
	"github.com/mingkid/jtt808-gateway/server/web/common"
	"github.com/mingkid/jtt808-gateway/server/web/common/errcode"
	"net/http"
)

type TerminalAPI struct{}

// Post 请求；新增终端
func (api TerminalAPI) Post(c *gin.Context) {
	type parameter struct {
		SN  string `json:"sn" binding:"required"`
		SIM string `json:"sim"`
	}
	var params parameter
	if err := c.ShouldBindJSON(&params); err != nil {
		common.NewErrorResponse(c, errcode.ParamsException.SetMsg(err.Error())).Return(http.StatusBadRequest)
		return
	}
	// 创建终端service
	termService := service.NewTerminal()
	err := termService.Save(service.TerminalSaveCommander{
		SN:  params.SN,
		SIM: params.SIM,
	})
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
