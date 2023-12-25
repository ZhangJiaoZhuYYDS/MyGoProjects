// @Author zhangjiaozhu 2023/3/17 18:01:00
package service

import (
	"Study/config/mysql"
	"Study/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


// GetRankList 问题相关接口
// @Summary 排名列表接口
// @Description 可按社区按分页，关键词或分类标识查询问题列表接口
// @Tags 问题相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity	query string false "category_id"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /ping [get]
func GetRankList(c *gin.Context)  {
	page , err:= strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Println("models problem Test page转换失败")
		return
	}
	size , _ := strconv.Atoi( c.DefaultQuery("size", "25"))

	page = (page -1) *size

	var count int64

	list :=make([]*models.UserBasic,0)

	err = mysql.DB.Model(new(models.UserBasic)).Count(&count).Order("finish_problem_num DESC , submit_num ASC").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg": "get rank_list err:"+err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code" :200 ,
		"data" :map[string]interface{}{
			"list":list,
			"count":count,
		},
	})
}
