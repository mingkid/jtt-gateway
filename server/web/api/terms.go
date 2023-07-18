package api

import (
	"github.com/gin-gonic/gin"
)

type TerminalsAPI struct{}

// Get 获取一组终端
func (t TerminalsAPI) Get(c *gin.Context) {

	return
}

func (t TerminalsAPI) Register(g *gin.RouterGroup) {
	r := g.Group("/terms")
	{
		r.GET("/registered", t.Get)
	}
}
