package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*声明单个变量*/
	var a int;
	fmt.Println(a)
	/*声明单个变量并赋值*/
	var b int = 2
	fmt.Println(b)
	var c = 29 // 可以推断类型
	fmt.Println(c)
	var d , e int = 1 , 2  // 声明多个变量
	fmt.Println(d,e)
	// 声明不同类型的变量
	var (
		f = 2
		g = "hello"
		h = 3.2
	)
	fmt.Println(f,g,h)
	// 简短声明( := 操作符的左边至少有一个变量是尚未声明的)
	i , j := 5 , 6 // 声明变量i和j
	j , k := 5 , 5 // j已经声明，但k尚未声明
	j , k = 10 , 11 // 给已经声明的变量j和k赋新值
	fmt.Println(i , j , k)
	fmt.Printf("type of k is %T , size of k is %d" , k ,unsafe.Sizeof(k))

	println(add(i, j))
	println(mianji(i , j))
	arr , _ := mianji(i , j)
	fmt.Println(arr)

}
func add(a , b int) int {
	c := a + b
	return c
}
//多返回值
func mianji(width , height int) (int , int)  {
	return width+height , width * height
}
//命名返回值
func mianji1(width , height int)(l , m int)  {
	l = width + height
	m = width * height
	return
}

