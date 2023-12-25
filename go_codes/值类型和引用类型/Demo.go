package main

import "fmt"

/*值类型和引用类型的说明
1) 值类型：基本数据类型int系列，float系列.bool,string、数组和结构体struct
	注意：变量直接存储值，内存通常在栈中分配
2) 引用类型：指针、slice切片、map、管道chan、interface等都是引用类型
	注意：变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存通常在
       堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾
       由GC来回收。
*/

func main() {
	//var a int = 6
	//var  b int  = 7
	//c := a++ //错误 a++ 只能独立使用
	//d := b--  //错误
	//a++   //正确
	var a string
	var b byte
	var c float32
	var d bool
	fmt.Println("请输入名字 年龄 薪水 是否")
	fmt.Scanf("%s %d %f %t" ,&a ,&b ,&c ,&d)
	fmt.Printf("名字--%v\n 年龄--%v\n 薪水--%v\n 是否--%v\n",a , b , c , d)


}