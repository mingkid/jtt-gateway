package common

import (
	"github.com/gin-gonic/gin"
	"github.com/mingkid/jtt-gateway/pkg/errcode"
)

type Response interface {
	Return(httpStatus int)
}

type SingleResponse struct {
	ctx  *gin.Context
	data any
}

func (resp SingleResponse) Return(httpStatus int) {
	if resp.data == nil {
		resp.data = gin.H{}
	}
	resp.ctx.JSON(httpStatus, gin.H{
		"errCode": 0,
		"errMsg":  "OK",
		"data":    resp.data,
	})
}

func NewSingleResponse(ctx *gin.Context, data any) SingleResponse {
	return SingleResponse{
		ctx:  ctx,
		data: data,
	}
}

type ListResponse struct {
	ctx        *gin.Context
	data       []any
	totalCount int64
}

func (resp ListResponse) Return(httpStatus int) {
	if resp.data == nil {
		resp.data = []any{}
	}
	resp.ctx.JSON(httpStatus, gin.H{
		"data": gin.H{
			"data": resp.data,
			"pagination": Pager{
				Page:      GetPage(resp.ctx),
				PageSize:  GetPageSize(resp.ctx),
				TotalRows: resp.totalCount,
			},
		},
		"errCode": 0,
		"errMsg":  "OK",
	})
}

func NewListResponse(ctx *gin.Context, data []any, totalCount int64) ListResponse {
	return ListResponse{
		ctx:        ctx,
		data:       data,
		totalCount: totalCount,
	}
}

type Pager struct {
	Page      uint  `json:"page"`     // 页码
	PageSize  uint  `json:"pageSize"` // 每页数量
	TotalRows int64 `json:"total"`    // 总行数
}

// NoPaginateListResponse 无分页列表
type NoPaginateListResponse struct {
	ctx  *gin.Context
	data []any
}

func (resp NoPaginateListResponse) Return(httpStatus int) {
	if resp.data == nil {
		resp.data = []any{}
	}
	resp.ctx.JSON(httpStatus, gin.H{
		"data":    resp.data,
		"errCode": 0,
		"errMsg":  "OK",
	})
}

func NoPaginatedListResponse(ctx *gin.Context, data []any) NoPaginateListResponse {
	return NoPaginateListResponse{
		ctx:  ctx,
		data: data,
	}
}

// ErrorResponse 错误异常返回
type ErrorResponse struct {
	ctx *gin.Context
	e   errcode.Error
}

func NewErrorResponse(ctx *gin.Context, e errcode.Error) ErrorResponse {
	// 临时用于事务
	return ErrorResponse{
		ctx: ctx,
		e:   e,
	}
}

func (resp ErrorResponse) Return(httpStatus int) {
	response := gin.H{
		"errCode": resp.e.Code(),
		"errMsg":  resp.e.Msg(),
		"data":    nil,
	}

	resp.ctx.JSON(httpStatus, response)
}
