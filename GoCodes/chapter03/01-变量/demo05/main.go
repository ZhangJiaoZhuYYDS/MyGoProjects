package main

import (
	"fmt"
	"unsafe"
)

//演示golang中+的使用
func main() {
	
	var i = 1
	var j = 2
	var r = i + j //做加法运算
	fmt.Println("r=", r)

	var str1 = "hello "
	var str2 = "world"
	var res = str1 + str2 //做拼接操作
	fmt.Println("res=", res)
	a := [3]int{1,2,3}
	b := []int{1,2,3}

	fmt.Printf("%T\n",&str1)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",&b)
	fmt.Printf("%T\n",main)
	var i8 int8
	fmt.Printf("%dB",unsafe.Sizeof(i8))
	fmt.Printf("%T",unsafe.Sizeof(i8))
	arr := 666
	fmt.Printf("%b\n",arr)
	fmt.Printf("%o\n",arr)
	fmt.Printf("%x\n",arr)
	ar := 3.6999
	fmt.Printf("%.0f",ar)

}