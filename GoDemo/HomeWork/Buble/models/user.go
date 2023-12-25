// @Author zhangjiaozhu 2023/3/4 20:47:00
package models

import (
	"encoding/json"
	"errors"
)

//  RegisterForm
//  @Description: 前端注册表信息

type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// RegisterForm
// @Description: 前端登录注册表
type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 后端数据表信息
type User struct {
	UserId   int64
	UserName string
	Password string
}

type VoteDataForm struct {
	//UserID int 从请求中获取当前的用户
	PostID    string  `json:"post_id" binding:"required"`	 // 帖子id
	Direction int8     `json:"direction,string" binding:"oneof=1 0 -1"`  // 赞成票(1)还是反对票(-1)取消投票(0)
}
// UnmarshalJSON 为VoteDataForm类型实现自定义的UnmarshalJSON方法
func (v *VoteDataForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction int8 `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}