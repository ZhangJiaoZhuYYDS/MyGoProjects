// @Author zhangjiaozhu 2023/3/15 17:40:00
package service

import (
	"Study/config/mysql"
	"Study/config/redis"
	"Study/models"
	"Study/utils/jwt"
	md52 "Study/utils/md5"
	"Study/utils/sendEmail"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


// GetUserDetail
// @Summary 用户详情
// @Description 查询用户的信息
// @Tags 用户方法
// @Param identity	query string false "user identity"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context)  {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1,
			"msg" : "用户唯一标识不能为空",
		})
		return
	}
	var data models.UserBasic
	err :=mysql.DB.Omit("password").Where("identity = ?" , identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1,
			"msg":"获取用户详情失败"+identity+"err:"+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":data,
	})
}

// Login
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户方法
// @Param username formData string false "user username"
// @Param password formData string false "user password"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /login [post]
func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1 ,
			"msg" : "用户密码不能为空",
		})
		return
	}
	password = md52.GetMd5(password)

	data := new(models.UserBasic)
	err := mysql.DB.Debug().Where("username = ? And password = ?",username,password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,gin.H{
				"code" : -1 ,
				"msg" : "用户密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}

	// 发送token
	token, err := jwt.GenerateToken(data.Identity, data.UserName, 0)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1,
			"msg" : "token err:"+err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"data": map[string]interface{}{
			"token":token,
		},
	})
}

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	checkCode := GetRand()
	// 将验证码存入redis,设置有效期
	redis.RDB.Set(c,email,checkCode,time.Second*300)
	// 将验证码发送到指定邮箱
	err := sendEmail.SendEmail(email,checkCode)
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

//生成uuid (字符串 36位)
func GetUUID()  string{
	return uuid.NewV4().String()
}
// 生成随机验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0 ; i < 6 ;i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

//注册

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param mail formData string true "mail"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if mail == "" || userCode == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	// 验证验证码是否正确
	sysCode, err := redis.RDB.Get(c, mail).Result()
	if err != nil {
		log.Printf("Get Code Error:%v \n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确，请重新获取验证码",
		})
		return
	}
	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	// 判断邮箱是否已存在
	var cnt int64
	err = mysql.DB.Where("mail = ?", mail).Model(new(models.UserBasic)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Error:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已被注册",
		})
		return
	}

	// 数据的插入
	userIdentity := GetUUID()
	data := &models.UserBasic{
		Identity: userIdentity,
		UserName:     name,
		Password: md52.GetMd5(password),
		Phone:    phone,
		Mail:     mail,
	}
	err = mysql.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Crete User Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"注册成功",
	})

}