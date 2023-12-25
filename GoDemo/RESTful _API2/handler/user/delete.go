package user

import (
	. "RESTful__API/handler"
	"RESTful__API/model"
	"RESTful__API/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier. 逻辑删除，只是修改删除时间
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
