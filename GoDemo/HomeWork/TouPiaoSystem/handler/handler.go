package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一一处理返回信息
type Reponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendReponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusNotImplemented, Reponse{
			Code:    100,
			Message: "错误",
			Data:    data,
		})
		return
	}
	c.JSON(http.StatusOK, Reponse{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}
