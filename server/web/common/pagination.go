package common

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt808-gateway/conf"
)

func GetPage(c *gin.Context) uint {
	type parameter struct {
		Page uint `form:"page"`
	}
	var params parameter
	if err := c.ShouldBindQuery(&params); err != nil {
		params.Page = 1
	}
	if params.Page <= 0 {
		params.Page = 1
	}
	return params.Page
}

func GetPageSize(c *gin.Context) uint {
	type parameter struct {
		PageSize uint `form:"page_size"`
	}
	var params parameter
	if err := c.ShouldBindQuery(&params); err != nil {
		params.PageSize = conf.Web.MaxPageSize
	}
	if params.PageSize <= 0 {
		params.PageSize = conf.Web.MaxPageSize
	}
	return params.PageSize
}
