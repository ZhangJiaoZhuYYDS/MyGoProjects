package user

import (
	"RESTful__API/handler"
	"RESTful__API/model"
	"RESTful__API/pkg/errno"
	"RESTful__API/pkg/token"
	"github.com/gin-gonic/gin"
)

// 用户登录监测 JWT
func Login(c *gin.Context) {
	// 绑定数据和结构体
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 根据请求的用户名获取数据库的用户信息
	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// 比较登录密码和数据库密码
	//if err := auth.Compare(d.Password, u.Password); err != nil {
	//	handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
	//	return
	//}
	if d.Password != u.Password {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	// 生成JWT并返回给客户端
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}
	handler.SendResponse(c, nil, model.Token{Token: t})
}
