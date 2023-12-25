// @Author zhangjiaozhu 2023/3/5 13:35:00
package Jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2 // 过期时间（两小时）
// 加密密钥
var mySecret = []byte("加密key")

// GetToken 生成JWT
func GetToken(userId int64, username string) (aToken string, err error) {
	// 创建一个自定义的数据
	c := MyClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 设置过期时间
			Issuer:    "zhang",                                    // 签发人
		},
	}
	//  使用指定的签名方法创建签名对象 (加密方式)   长token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	// 短期token
	//rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
	//	Issuer:    "zhang",                                 // 签发人
	//}).SignedString(mySecret)
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	//校验token
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token(无效令牌)")
}
