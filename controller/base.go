package controller

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/model/response"
)

type TableResponse struct {
	Data       interface{} `json:"data""`
	PageSize   int         `json:"pageSize""`
	PageNo     int         `json:"pageNo"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
}

func ResponseTableSuccess(c *gin.Context, data interface{}, pageSize, pageNo, totalPage, totalCount int, msg string) {
	var tableResponse TableResponse
	tableResponse.PageSize = pageSize
	tableResponse.PageNo = pageNo
	tableResponse.TotalPage = totalPage
	tableResponse.TotalCount = totalCount
	tableResponse.Data = data
	var base response.Base
	base.Status = 200
	base.Result = &tableResponse
	base.Message = msg
	c.JSON(200, base)
}
func ResponseSuccess(c *gin.Context, data interface{}, msg string) {
	var base response.Base
	base.Status = 200
	base.Result = data
	base.Message = msg
	c.JSON(200, base)
}
func ResponseFailed(c *gin.Context, data interface{}, msg string) {
	var base response.Base
	base.Status = 500
	base.Result = data
	base.Message = msg
	c.JSON(200, base)
}
