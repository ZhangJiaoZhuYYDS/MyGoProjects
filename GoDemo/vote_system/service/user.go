// @Author zhangjiaozhu 2023/3/18 13:46:00
package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"vote_system/dao"
	"vote_system/models"
	"vote_system/utils/md5"
	"vote_system/utils/mysql"
	"vote_system/utils/redis"
	"vote_system/utils/sendEmail"
)
// Register
// @Tags 用户
// @Summary 用户注册
// @Param mail formData string true "mail"
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param sex formData string true "sex"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(c *gin.Context)  {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code" : 101,
			"msg" : "数据参数格式校验错误",
		})
		return
	}
	// 判断邮箱和用户名是否存在
	var count int64
	err = mysql.DB.Model(new(models.User)).Debug().Where("mail = ? or username = ?", user.Mail, user.Username).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"服务器繁忙",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"邮箱或用户名已存在",
		})
		return
	}
	err = dao.Save(user)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"数据库服务器繁忙",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200 ,
		"msg" : "注册成功",
	})
}


func SendEmailCode(c *gin.Context) {
	type VerificationRequest struct {
		Mail string `json:"mail"`
	}
	var req VerificationRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	checkCode := GetRand()
	// 将验证码存入redis,设置有效期
	redis.RDB.Set(c,req.Mail,checkCode,time.Second*300)
	// 将验证码发送到指定邮箱
	err = sendEmail.SendEmail(req.Mail,checkCode)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Code Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// SendEmailCode2
// @Tags 公共方法
// @Summary 使用邮箱发送验证码
// @Param mail formData string true "mail"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /sendemailcode [post]
func SendEmailCode2(c *gin.Context) {
	mail := c.PostForm("mail")
	type VerificationRequest struct {
		Mail string `json:"mail"`
	}
	var req VerificationRequest
	req.Mail = mail
	if req.Mail == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	checkCode := GetRand()
	// 将验证码存入redis,设置有效期
	redis.RDB.Set(c,req.Mail,checkCode,time.Second*3000)
	// 将验证码发送到指定邮箱
	err := sendEmail.SendEmail(req.Mail,checkCode)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Code Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}


// 生成6位随机验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0 ; i < 6 ; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// 登录

// Login
// @Summary 跨域用户登录
// @Produce json
// @Tags 跨域用户
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param code formData string true "验证码"
// @Success 200 {string} json "成功"
// @Router /login [post]
func Login(c *gin.Context)  {
	var req struct{
		Username string `json:"username"`
		Password string `json:"password"`
		Code string 	`json:"checkword"`
		Mail string `json:"mail"`
	}
	if err := c.BindJSON(&req) ; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"msg":"请求参数有误",
		})
		return
	}

	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//code := c.PostForm("code")
	// 基础校验
	if req.Username== "" || req.Password == "" || req.Code == "" || req.Mail == "" {
		c.JSON(http.StatusOK,gin.H{
			"code" : 101,
			"msg" : "参数不能为空",
		})
		return
	}
	// 验证验证码
	result, err := redis.RDB.Get(c, req.Mail).Result()
	if err != nil {
		log.Info("service Login redis err "+err.Error())
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"验证码不正确，请重新获取验证码",
		})
		return
	}
	if result != req.Code {
		c.JSON(http.StatusOK,gin.H{
			"code": 101,
			"msg":"验证码输入错误，请重新输入",
		})
		return
	}
	// 查询数据库
	user, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,gin.H{
				"code":101,
				"msg":"账号不存在",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"服务器繁忙",
		})
		return
	}
	// 校验加密密码
	if md5.GetMd5(req.Password) != user.Password {
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"密码错误",
		})
	}
	// 设置cookie ，session ,Jwt
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 3000})
	//设置 Session
	session.Set("UserId", user.UserIdentity)
	err = session.Save()
	if err != nil {
		log.Info("session err"+err.Error())
		return
	}
	// 跳转首页
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}

// Login2
// @Summary 用户登录
// @Produce json
// @Tags 用户
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param code formData string true "验证码"
// @Success 200 {string} json "成功"
// @Router /logins [post]
func Login2(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := c.PostForm("code")
	// 基础校验
	if username== "" || password == "" || code == ""  {
		c.JSON(http.StatusOK,gin.H{
			"code" : 101,
			"msg" : "参数不能为空",
		})
		return
	}
	// 验证验证码
	result, err := redis.RDB.Get(c, "2568317193@qq.com").Result()
	if err != nil {
		log.Info("service Login redis err "+err.Error())
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"验证码不正确，请重新获取验证码",
		})
		return
	}
	if result != code {
		c.JSON(http.StatusOK,gin.H{
			"code": 101,
			"msg":"验证码输入错误，请重新输入",
		})
		return
	}
	// 查询数据库
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,gin.H{
				"code":101,
				"msg":"账号不存在",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"服务器繁忙",
		})
		return
	}
	// 校验加密密码
	if md5.GetMd5(password) != user.Password {
		c.JSON(http.StatusOK,gin.H{
			"code":101,
			"msg":"密码错误",
		})
		return
	}
	// 设置cookie ，session ,Jwt
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 3600})
	//设置 Session
	session.Set("UserId", user.UserIdentity)
	err = session.Save()
	if err != nil {
		log.Info("session err"+err.Error())
		return
	}
	// 跳转首页
	c.JSON(200,gin.H{
		"code":200,
		"msg":"登录成功",
	})
}


// Logout
// @Summary 用户退出
// @Produce json
// @Tags 用户
// @Success 200 {string} json "成功"
// @Router /votes/logout [get]
func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Delete("UserId")
	//session.Clear()
	err := session.Save()
	if err != nil {
		log.Info("session err"+err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"退出成功",
	})
}