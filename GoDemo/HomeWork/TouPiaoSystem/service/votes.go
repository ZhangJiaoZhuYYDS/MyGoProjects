// @Author zhangjiaozhu 2023/3/9 9:06:00
package service

import (
	"TouPiaoSystem/dao"
	"TouPiaoSystem/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
	"time"
)

func CreateVote(c *gin.Context)  {
	var form model.VoteForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg" : "无效参数",
		})
		return
	}

	userId, exists := c.Get("UserId")
	if !exists {
		c.JSON(http.StatusOK,gin.H{
			"msg":"用户未登录",
		})
		return
	}


	var vote model.Votes

	vote.Title = form.Title
	vote.Type = form.Type



	// 判断投票单是否过期
	model.DB.Self.Where("type = ?",form.Type).First(&vote)

	if vote.Type == "" {
		c.JSON(http.StatusOK,gin.H{
			"msg":"不存在此种类型的投票",
		})
		return
	}
	start := vote.CreatedAt
	during ,err := strconv.Atoi(vote.Title)
	toUint64 := cast.ToInt64(during)
	if start.Unix() + (toUint64*60*60*1000) < time.Now().Unix() {
		c.JSON(http.StatusOK,gin.H{
			"msg":"无法投票，投票已过期",
		})
		return
	}
	// 开启事务 ，防止用户投过票后总票数重复计票的问题
	tx := model.DB.Self.Begin()

	defer func() {
		if r:= recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 查询是否存在投票箱
	var a model.Votes
	tx.Where("type = ?" , form.Type).First(&a)
	if a.Type != "" {
		// 更新投票项的票数
		// 投票数加一
		totalNumber := 0
		tx.Debug().Model(&model.Votes{}).Select("total_number").Find(&totalNumber)
		vote.TotalNumber = totalNumber+1

		err :=tx.Debug().Model(&model.Votes{}).Where("type = ?" , form.Type).Update("total_number",vote.TotalNumber).Error
		if err!= nil{
			c.JSON(http.StatusOK,gin.H{
				"msg" : "服务器内部错误",
			})
			return
		}
	}else {
		// 创建投票项
		// 创建单个类型的投票箱

		// 更新投票项的票数
		// 投票数加一
		totalNumber := 0
		tx.Debug().Model(&model.Votes{}).Select("total_number").Find(&totalNumber)
		vote.TotalNumber = totalNumber+1
		tx.Create(&vote)
	}


	// 判断是否刷票
	status := dao.FindUserStatus(cast.ToUint64(userId),form.Type)
	if status == 1 {
		// 事务回滚
		tx.Rollback()
		c.JSON(http.StatusOK,gin.H{
			"msg": "您已投过票，无法重复投票",
		})
		return
	}
	err = dao.CreateVote(form, cast.ToUint64(userId))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg":"服务器繁忙",
		})
		return
	}


	
	tx.Commit()


	c.JSON(http.StatusOK,gin.H{
		"data": form,
		"msg":"投票成功",
	})
}
//通过id获取用户投票信息
func UserVote(c *gin.Context)  {
	userId, exists := c.Get("UserId")
	if !exists {
		c.JSON(http.StatusOK,gin.H{
			"msg":"用户未登录",
		})
		return
	}
	userVotes := dao.GetVoteById(cast.ToUint64(userId))
	c.JSON(http.StatusOK,gin.H{
		"data":userVotes,
	})

}
func AllUserVote(c *gin.Context)  {
	all := dao.GetVoteAll()
	c.JSON(http.StatusOK,gin.H{
		"data":all,
	})
}


func TotalVote(c *gin.Context) {
	Votetype := c.Query("type")
	number := 0
	model.DB.Self.Select("type").Where("type = ?",Votetype).Find(&number)
	c.JSON(http.StatusOK,gin.H{
		"data":number,
	})
}

func VoteType(c *gin.Context)  {
	var votetype []string
	model.DB.Self.Model(&model.Votes{}).Select("type").Find(&votetype)
	c.JSON(http.StatusOK,gin.H{
		"data":votetype,
	})
}