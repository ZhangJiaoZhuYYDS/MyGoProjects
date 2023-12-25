package middlewares

// 记录请求响应前和响应后的动作

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 全局可访问  example 变量
		c.Set("example", "12345")
		fmt.Println("全局中间件，在所有请求前运行，并设置了全局变量‘example’,另外设置了开始时间，用于计时")
		// 请求前

		c.Next()
		fmt.Println("全局中间件，在请求结束后执行，记录了请求结束时间")
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println("获取响应状态码:", status)
	}
}
