package main

import "fmt"

func main() {
	var a int //声明实际变量
	var b *int //声明指针变量
	b = &a  // 1 指针变量的存储地址   &a
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b) //2 使用指针变量访问访问值  *b

	var c *int  //空指针：声明指针变量但是不赋值
	if c==nil {
		fmt.Println("这是空指针")
	}
	var f int = 6
	var g int = 9
	k := &f
	l := &g
	fmt.Println(k,l)
	k , l = l , k
	fmt.Println(k,l)
	fmt.Println(*k,*l)
	f , g = *k , *l

	*k = 100
	fmt.Println(f,g)

	change(k)
	fmt.Println("通过指针参数改变的值为",f)
}
// 3 函数传递指针参数
func change(a *int)  {
	*a = 55
}
