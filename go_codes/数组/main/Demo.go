package main

import "fmt"

func main() {
	var a [3]int  //数组的声明(默认值是0)
	b := [3]int{7,8,9}  //数组简略声明
	c := [...]string{"USA","CHINA","UK","FRANCE"}
	d := c  //数组复制
	d[0] = "JAPAN"
	a[0] = 1
	a[1] = 2
	fmt.Println("普通声明数组",a)
	fmt.Println("简易声明数组",b)
	fmt.Println("666")
	fmt.Println(c,d)
	fmt.Println("数组长度是len()",len(c))
	//遍历数组
	for i := 0 ; i < len(c) ; i++ {
		fmt.Printf("%v 是 %v\n" , i , c[i])
	}
	  //高级遍历
	sum := ""
	for i , v := range c {  // i 是索引 v 是元素
		fmt.Println(i,"-----",v)
		sum = sum + v
	}
	fmt.Println(sum)
}
