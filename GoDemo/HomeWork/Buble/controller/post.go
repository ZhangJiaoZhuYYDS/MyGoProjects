// @Author zhangjiaozhu 2023/3/7 19:21:00
package controller

import (
	"Buble/models"
	"Buble/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func CreatePostController(c *gin.Context)  {
	// 1 获取参数及校验参数
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		log.Info("校验参数失败")
		ResponseErrorWithMsg(c,models.CodeInvalidParams,err.Error())
		return
	}
	// 参数校验
	// 获取作者id(从请求上下文中获取用户id) (token检验成功后，会解析token获取userid,并存储在请求上下文中)
	id, err := getCurrentUserID(c)
	if err != nil {
		log.Info("获取上下文中的id失败")
		ResponseErrorWithMsg(c,models.CodeNotLogin,err.Error())
		return
	}
	post.AuthorID = id
	// 2 创建帖子
	err = service.CreatePost(&post)
	if err != nil {
		log.Info("创建帖子失败")
		ResponseError(c,models.CodeServerBusy)
		return
	}
	// 3
	ResponseSuccess(c , nil)

}

func PostListsController(c *gin.Context)  {
	 getPageInfo(c)
}
