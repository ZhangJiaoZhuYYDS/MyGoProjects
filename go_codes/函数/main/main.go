package main

import (
	"fmt"
	"go_codes/函数/utils"
)

func main() {
	fmt.Println(utils.Name) //引入外部包中的属性
	utils.Method() //引入外部包中的方法
	di(4)
	fmt.Println(method(6))

}
func init() {
	fmt.Println("main包初始化")
}
func method(n int) (r int) {
	r = n
	return
}
func out(a int)  {
	if a > 4 {
		return
	}
	a++
	out(a)
}
// 递归函数
func di( a int)  {
	if a > 2 {
		a--
		di(a)
	}
	fmt.Println("a是",a)
}
