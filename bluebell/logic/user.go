// @Author zhangjiaozhu 2023/6/4 10:24:00
package logic

import (
	"bluebell/common/snowflake"
	"bluebell/dao/mysql"
	"bluebell/models"
	"errors"
)

func SignUp(p *models.ParamSignUp) error {
	// 1 判断用户是否存在
	name, err := mysql.QueryUserByName(p.Username)
	if err != nil {
		return nil
	}
	if name {
		return errors.New("用户已存在")
	}
	// 2 生成UID
	id := snowflake.GenId()
	u := &models.User{
		UserId:   id,
		Username: p.Username,
		Password: p.Password,
	}
	// 3 保存数据库
	err = mysql.InsertUser(u)
	if err != nil {
		return err
	}
	return nil
}

func Login(u *models.ParamLogin) error {
	user := &models.User{
		Username: u.Username,
		Password: u.Password,
	}
	err := mysql.Login(user)
	return err
}
