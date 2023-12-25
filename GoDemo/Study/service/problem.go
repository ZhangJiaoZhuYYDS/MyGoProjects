// @Author zhangjiaozhu 2023/3/11 22:40:00
package service

import (
	"Study/config/mysql"
	"Study/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList 问题相关接口
// @Summary 问题列表接口
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
func GetProblemList(c *gin.Context)  {
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	page , err:= strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Println("models problem Test page转换失败")
		return
	}
	size , _ := strconv.Atoi( c.DefaultQuery("size", "25"))

	page = (page -1) *size

	var count int64

	list := make([]*models.ProblemBasic,0)
	tx := models.GetProblemList(keyword,categoryIdentity)
	err = tx.Debug().Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("modes problem Test 查询数据库失败")
		return
	}

	c.JSON(200,gin.H{
		"msg" : "6666",
		"data" : map[string]interface{}{
			"count":count,
			"list":list,
		},
	})
}


// GetProblemDetail 问题相关接口
// @Summary 问题详情
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 问题相关接口
// @Param identity	query string false "problem identity"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context)  {
	identity :=c.Query("identity")
	if identity == "" {
		c.JSON(200,gin.H{
			"code" : -1 ,
			"msg" : "问题标识不能为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := mysql.DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,gin.H{
				"code" : -1,
				"msg" : "查询记录不存在",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"code": -1 ,
			"msg" : "get problem_detail error:" +err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : 200 ,
		"data" :data,
	})
}