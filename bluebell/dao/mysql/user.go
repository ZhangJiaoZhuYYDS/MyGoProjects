// @Author zhangjiaozhu 2023/6/4 10:26:00
package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
)

func InsertUser(user *models.User) error {
	sqlStr := `insert into user(user_id, username, password) values (?,?,?) `
	_, err := db.Exec(sqlStr, user.UserId, user.Username, EncryptPassword(user.Password))
	if err != nil {
		return err
	}
	return nil
}
func QueryUserById(id int64) {

}
func QueryUserByName(s string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, s); err != nil {
		return false, err
	}
	return count > 0, nil
}
func Login(u *models.User) error {
	opassword := u.Password
	sqlStr := `select user_id , username , password from user where username = ?`
	err := db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	if EncryptPassword(opassword) != u.Password {
		return errors.New("密码错误")
	}
	return nil
}
