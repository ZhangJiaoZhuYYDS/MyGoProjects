// @Author zhangjiaozhu 2023/3/7 17:09:00
package controller

import (
	"Buble/models"
	"Buble/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CommunityHandler
//  @Description: 查询服务社区列表
//  @Author zhangjiaozhu 2023-03-07 18:01:01
//  @param c
//
func CommunityHandler(c *gin.Context) {
	list, err := service.GetCommunityList()
	if err != nil {
		fmt.Println("服务层查询失败")
		ResponseError(c,models.CodeServerBusy)
	}
	ResponseSuccess(c,list)
}
// CommunityHandlerDetail
//  @Description: 查询单个服务详情
//  @Author zhangjiaozhu 2023-03-07 18:01:17
//  @param c
//
func CommunityHandlerDetail(c *gin.Context)  {
	// 1 获取id
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ResponseError(c,models.CodeInvalidParams)
		return
	}
	// 2 根据id查询社区详情
	byId, err := service.GetCommunityById(id)
	if err != nil {
		ResponseErrorWithMsg(c,models.CodeSuccess,err.Error())
		return
	}
	ResponseSuccess(c,byId)
}