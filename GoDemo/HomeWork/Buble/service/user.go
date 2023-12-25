// @Author zhangjiaozhu 2023/3/4 21:04:00
package service

import (
	"Buble/commons/Jwt"
	"Buble/commons/snowflake"
	"Buble/dao/mysql"
	"Buble/models"
	"errors"
)

// Register
//
//	@Description:
//	@Author zhangjiaozhu 2023-03-05 09:05:25
//	@param p
//	@return error
func Register(p *models.RegisterForm) (error error) {
	// 1 判断用户是否存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		// 数据库查询出错
		return errors.New("查询数据库失败")
	}
	// 2 生成UUID
	userid := snowflake.GenID()
	user := models.User{
		UserId:   userid,
		UserName: p.UserName,
		Password: p.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.LoginForm) (atoken string, error error) {
	user := models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(&user); err != nil {
		return "", err
	}
	// 登录成功，生成JWT
	return Jwt.GetToken(user.UserId, user.UserName)
}
