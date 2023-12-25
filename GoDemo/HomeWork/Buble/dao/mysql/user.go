// @Author zhangjiaozhu 2023/3/4 21:07:00
package mysql

import (
	"Buble/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "zhangjiaozhu.vip"

/**
 * @Author zhangjiaozhu
 * @Description //TODO 对密码进行加密
 * @Date 21:50 2023/03/01
 **/
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

// CheckUserExist
//
//	@Description: 检查用户名是否存在
//	@Author zhangjiaozhu 2023-03-05 11:07:47
//	@param username
//	@return error
func CheckUserExist(username string) error {
	var count int
	err := _db.Debug().Raw("select count(user_id) from user where username = ?", username).Scan(&count).Error
	if err != nil {
		return errors.New("查询数据库失败")
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}
	return nil
}

func InsertUser(user models.User) error {
	// 对密码加密
	user.Password = encryptPassword([]byte(user.Password))
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	err := _db.Debug().Exec(sqlStr, user.UserId, user.UserName, user.Password).Error
	return err
}

func Login(user *models.User) (error error) {
	originPassword := user.Password // 记录用户的原始登录密码
	sqlStr := `select user_id ,username,password from user where username = ?`
	err := _db.Debug().Raw(sqlStr, user.UserName).Scan(user).Error
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return errors.New("查询数据库出错")
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return errors.New("用户不存在")
	}
	// 生成加密密码与数据库密码进行比较
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return errors.New("密码错误")
	}
	return
}
