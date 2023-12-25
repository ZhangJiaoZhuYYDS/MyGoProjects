package user

import (
	"TouPiaoSystem/handler"
	"TouPiaoSystem/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 登录校验
func Login(c *gin.Context) {
	var user model.UserModel
	// 把表单json形式请求数据和结构体数据双相绑定
	err := c.ShouldBind(&user)
	if err != nil {
		handler.SendReponse(c, err, "请求数据格式不匹配")
		return
	}
	// 跳转到投票网页
	//c.Redirect(http.StatusFound, "/index")
	// 校验数据库账号和密码
	d, err := model.GetUser(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": " 101",
			"msg":  "账号不存在",
		})
		return
	}
	if user.Password != d.Password {
		// 直接返回投票页面
		c.JSON(http.StatusOK, gin.H{
			"code": "102",
			"msg":  "密码错误",
		})
		return
	}
	// 设置cookie
	cookie, err := c.Cookie("login_cookie")
	if err != nil {
		cookie = "no set"
		// cookie 的值是用户的id
		c.SetCookie("login_cookie", strconv.FormatUint(d.Id, 10), 3600, "/", "mydemo.com:8090", false, true)
	}
	fmt.Println(cookie)
	c.HTML(http.StatusOK, "form.html", gin.H{
		"code": "100",
	})
}
