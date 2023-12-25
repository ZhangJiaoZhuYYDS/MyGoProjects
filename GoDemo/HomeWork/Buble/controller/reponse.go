// @Author zhangjiaozhu 2023/3/7 17:41:00
package controller

import (
	"Buble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 封装统一响应数据

type ResponseData struct {
	Code models.Code 	`json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"date,omitempty"` // omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, c models.Code) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code models.Code, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    models.CodeSuccess,
		Message: models.CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}