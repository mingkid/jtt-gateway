package internal

import (
	"github.com/gin-gonic/gin"
)

type TerminalsAPI struct{}

// get 获取一组终端
func (t TerminalsAPI) get(c *gin.Context) {

	return
}

func (t TerminalsAPI) Register(g *gin.RouterGroup) {
	r := g.Group("/terms")
	{
		r.GET("/registered", t.get)
	}
}
