package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// 分发不同请求

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		//检查请求头是否存在以下字段
		requestId := c.Request.Header.Get("X-Request-Id")
		//若没有，就新建一个
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}
		//设置请求头的id
		c.Set("X-Request-Id", requestId)
		//
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
