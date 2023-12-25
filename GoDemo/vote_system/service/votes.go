// @Author zhangjiaozhu 2023/3/19 17:15:00
package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vote_system/dao"
	"vote_system/models"
	"vote_system/utils/mysql"
)


// DoVote
// @Summary 用户投票
// @Produce json
// @Tags 投票
// @Param result formData string true "是否赞成 0 赞成 1 反对"
// @Param hobbies formData string false "你的爱好，可以填：html,css,java,go,c++"
// @Param type formData string true "投票类型，比如：赞成反对票"
// @Success 200 {string} json "成功"
// @Router /votes/doVote [post]
func DoVote(c *gin.Context)  {


	type postform struct {
		Result string	`json:"result" gorm:"result"`  // 是否赞成 0 赞成  1 反对
		Hobbies string `json:"hobbies" gorm:"column:hobbies"`  // 兴趣爱好
		Status int 	`gorm:"column:status;default:0"`   // 投票状态 0 未投票 1 已投票
		Type string `gorm:"column:type;"` // 投票类型
	}
	var p postform
	err2 := c.ShouldBind(&p)
	// 参数校验
	if err2 != nil {
		c.JSON(200,gin.H{
			"code":101,
			"msg":"请求参数有误",
		})
		return
	}
	if p.Type == "" || p.Result == "" {
		c.JSON(200,gin.H{
			"code":101,
			"msg":"投票项不能为空",
		})
		return
	}
	// 绑定结构体
	var Uservotes models.Votesopt
	Uservotes.Result, _ = strconv.Atoi(p.Result)
	Uservotes.Type = p.Type
	Uservotes.Hobbies = p.Hobbies
	// 获取用户id
	value, exists := c.Get("user_id")
	if !exists {
		c.JSON(200,gin.H{
			"code":101,
			"msg":"服务器繁忙",
		})
		return
	}
	s , ok := value.(string)
	if !ok {
		fmt.Println("any 转 string err")
		c.JSON(200,gin.H{
			"code":101,
			"msg":"服务器繁忙",
		})
	}


	// 判断票单是否存在

	var votes models.Votes
	err := mysql.DB.Where("type = ?", p.Type).First(&votes).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println("数据库查询失败")
			return
		}
	}
	if votes.Type == "" {
		Uservotes.Status = 1
		Uservotes.UserIdentity = s
		err := mysql.DB.Create(&Uservotes).Error
		if err != nil {
			c.JSON(200,gin.H{
				"code":101,
				"msg":"创建用户表单失败",
			})
		}
		votes.Type = p.Type
		votes.TotalNumber = 1
		mysql.DB.Create(&votes)
	}else {
		// 获取票箱的创建时间和持续时间
		start := votes.CreatedAt
		//during :=votes.During
		//toUint64 := cast.ToInt32(during)
		//if start.Unix() + (toUint64*60*60*1000) < time.Now().Unix() {
		if start.Add(time.Hour*1).Unix() < time.Now().Unix() {
			c.JSON(http.StatusOK,gin.H{
				"msg":"无法投票，投票已过期",
			})
			return
		}
		// 判断当前用户是否刷票
		status,err := dao.FindUserStatus(s,p.Type)
		if err != nil {
			// 根据userIdentity查询用户信息不存在，就新建一个新用户投票信息，总票数+1
			if err == gorm.ErrRecordNotFound {
				Uservotes.Status = 1
				Uservotes.UserIdentity = s
				err := mysql.DB.Create(&Uservotes).Error
				if err != nil {
					c.JSON(200,gin.H{
						"code":101,
						"msg":"创建用户表单失败",
					})
					return
				}
				votes.TotalNumber+= 1
				mysql.DB.Model(&models.Votes{}).Where("type = ?",p.Type).Update("total_number",votes.TotalNumber)
				c.JSON(200,gin.H{
					"code":200,
					"msg":"success",
				})
				return

			}
			log.Println("查询刷票失败")
			c.JSON(200,gin.H{
				"code":101,
				"msg":"数据库查询失败",
			})
			return
		}
		if status == 1 {
			// 事务回滚
			c.JSON(http.StatusOK,gin.H{
				"msg": "您已投过票，无法重复投票",
			})
			return
		}
		if status == 0 {
			mysql.DB.Model(&models.Votesopt{}).Where("user_identity = ?",s).Update("status",1)
			votes.TotalNumber = votes.TotalNumber+1
			mysql.DB.Model(&models.Votes{}).Where("type = ?",p.Type).Update("total_number",votes.TotalNumber)
		}


	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})

}

// GetVotesByType
// @Summary 查询某类型投票信息
// @Produce json
// @Tags 投票
// @Param type_1 formData string true "投票类型，"
// @Success 200 {string} json "成功"
// @Router /votes/getVotesByType [post]
func GetVotesByType(c *gin.Context)  {
	type S struct {
		Type string `json:"type"`
	}
	var s S
	err2 := c.ShouldBind(&s)
	if err2 != nil {
		c.JSON(200,gin.H{
			"code":101,
			"msg":"格式错误",
		})
		return
	}
	votes, count,err := dao.GetVotesByType(s.Type)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(200,gin.H{
				"code":101,
				"msg":"此种类型的投票不存在",
			})
			return
		}
		c.JSON(200,gin.H{
			"code":101,
			"msg":"数据库服务器繁忙",
		})
	}
	a := 0
	b := 0
	for i := range votes.VotesOpt {
		if votes.VotesOpt[i].Result==0 {
			a++
		}
		if votes.VotesOpt[i].Result==1 {
			b++
		}
	}
	c.JSON(200,map[string]interface{}{
		"code":200,
		"count":count,
		"data": map[string]int{
			"0":a,
			"1":b,
		},
	})
}

// GetVotesoptByType
// @Summary 查询当前用户的投票信息
// @Produce json
// @Tags 投票
// @Param type_1 formData string true "投票类型"
// @Success 200 {string} json "成功"
// @Router /votes/getVotesoptByType [post]
func GetVotesoptByType(c *gin.Context)  {
	session :=sessions.Default(c)
	votesType := c.PostForm("type_1")
	userIdentity := session.Get("UserId")
	votesopt, err := dao.GetVotesoptByType(userIdentity.(string), votesType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(200,gin.H{
				"code":101,
				"msg":"查不到用户此种类型的投票记录",
			})
			return
		}
		c.JSON(200,gin.H{
			"code":101,
			"msg":"数据库繁忙",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":votesopt,
	})
}
// GetVotes
// @Summary 查询所有种类的投票信息
// @Produce json
// @Tags 投票
// @Success 200 {string} json "成功"
// @Router /votes/getVotes [get]
func GetVotes(c *gin.Context)  {
	var votes []models.Votes
	mysql.DB.Model(models.Votes{}).Find(&votes)
	c.JSON(200,gin.H{
		"data":votes,
	})
}

// CreateVote 创建投票
// @Summary 用户创建投票
// @Produce json
// @Tags 投票
// @Param title formData string true "投票的标题:比如：你赞成吃饭吗？"
// @Param during formData string true "投票的有效期，以小时为单位"
// @Param opt formData string true "投票的内容：比如：result,hobbies,以空格隔开"
// @Param type formData string true "投票类型，比如：type_1"
// @Success 200 {string} json "成功"
// @Router /votes/createVote [post]
func CreateVote (c *gin.Context){
	titleVote := c.PostForm("title")
	duringVote := c.PostForm("during")
	optVote := c.PostForm("opt")
	typeVote := c.PostForm("type")

	var myVote models.Votes
	myVote.Type = typeVote
	myVote.During, _ = strconv.Atoi(duringVote)
	myVote.Title = titleVote
	myVote.During = 1

	opts :=strings.Split(optVote," ")
	fmt.Println(opts[:])
	result := dao.NewVote(&myVote)
	if result {
		c.JSON(200,gin.H{
			"code":200,
			"msg":"创建成功",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":101,
		"msg":"创建失败",
	})
}