// @Author zhangjiaozhu 2023/3/7 21:02:00
package controller

import (
	"Buble/models"
	"Buble/service"
	"github.com/gin-gonic/gin"
)

func VoteController(c *gin.Context)  {
	// 参数校验，给哪个文章投了什么票
	vote := new(models.VoteDataForm)
	err := c.ShouldBindJSON(&vote)
	if err != nil {
		ResponseSuccess(c,"参数绑定失败")
		return
	}
	// 获取当前请求用户的id
	id, err := getCurrentUserID(c)
	if err != nil {
		ResponseSuccess(c,models.CodeNotLogin)
		return
	}
	// 投票逻辑
	service.VoteForPost(id,vote)
}
