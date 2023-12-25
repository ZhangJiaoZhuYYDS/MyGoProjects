package user

import (
	"RESTful__API/handler"
	"RESTful__API/model"
	"RESTful__API/pkg/errno"
	"RESTful__API/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"net/http"
)

func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	var err error
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	fmt.Printf("username is: [%s], password is [%s]", r.Username, r.Password)
	//
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 校验用户名
	//if r.Username == "" {
	//	err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
	//	log.Errorf(err, "Get an error")
	//}
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 校验用户密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
	}

	// 保存密码到数据库
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
	}
	rsp := CreateResponse{Username: r.Username}

	handler.SendResponse(c, nil, rsp)

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
