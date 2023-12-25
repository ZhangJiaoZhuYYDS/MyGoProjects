// @Author zhangjiaozhu 2023/3/17 10:49:00
package test

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUUID(t *testing.T)  {
	s := uuid.NewV4().String()
	println(s)
}
