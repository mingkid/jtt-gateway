package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt-gateway/domain/service"
	"github.com/mingkid/jtt-gateway/http/internal/common"
	"github.com/mingkid/jtt-gateway/pkg/errcode"
)

// PlatformAuth 平台访问认证
func PlatformAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var args accessToken
		if err := c.ShouldBindQuery(&args); err != nil {
			common.NewErrorResponse(c, errcode.ParamsError.SetMsg(err.Error())).Return(http.StatusBadRequest)
			return
		}

		var pSvr service.Platform
		_, err := pSvr.GetByToken(args.Token)
		if err != nil {
			common.NewErrorResponse(c, errcode.UnauthorizedError).Return(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

type accessToken struct {
	Token string `form:"token" binding:"required"` // 访问令牌
}
