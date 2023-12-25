// @Author zhangjiaozhu 2023/3/16 16:55:00
package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type UserClaim struct {
	Identity string `json:"identity"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var mykey = []byte("密钥")
// 生成token
func TestGenerateToken( t *testing.T)  {
	userClaim := UserClaim{
		Identity:       "user_1",
		Name:           "Get",
		StandardClaims: jwt.StandardClaims{},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,userClaim)
	tokenString, err := token.SignedString(mykey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tokenString)
}
// 解析token
func TestAnalyseToken(t *testing.T)  {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiJHZXQifQ.fd9pJEf0Du_2Lfp2wRdKszrvzo_vuprGHJ--gIJDxo8"
	useClaim := new(UserClaim)
	claim , err :=jwt.ParseWithClaims(tokenString,useClaim, func(token *jwt.Token) (interface{}, error) {
		return mykey , nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claim.Valid {
		fmt.Println(useClaim)
	}
}