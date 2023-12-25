// @Author zhangjiaozhu 2023/3/15 22:47:00
package service

import (
	"Study/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Tags 提交记录相关接口
// @Summary 提交记录列表接口
// @Description 可按社区按分页，关键词或分类标识查询问题列表接口
// @Accept application/json
// @Produce application/json
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Param status query int false "status 0待判断 1正确 2 错误 3 超时"
// @Param user_identity query string false "user_identity"
// @Param problem_identity	query string false "problem_identity"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /ping [get]
func GetSubmitList(c *gin.Context)  {
	page , err:= strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Println("models problem Test page转换失败")
		return
	}

	size , _ := strconv.Atoi( c.DefaultQuery("size", "25"))
	page = (page -1) *size

	var count int64
	list := make([]models.SubmitBasic,0)

	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("problem_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	tx :=models.GetSubmitList(problemIdentity,userIdentity,status)

	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1 ,
			"msg" : "get submit list error :" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : 200 ,
		"data" : map[string]interface{}{
			"list" : list,
			"count" : count,
		},
	})
}
