package model

// 用户信息模型
type UserModel struct {
	BaseModel
	Username string `json:"username" form:"username" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

// 根据用户名查询单个用户信息
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Debug().Where("username = ? ", username).First(&u)
	return u, d.Error
}
