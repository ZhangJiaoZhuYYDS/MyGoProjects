package main

import (
	"fmt"
)

func main() {
	if a := 6 ; a > 5 {
		fmt.Println("特殊用法")
	}
	if a := 7 ; a > 6 {
		fmt.Println("66666666")
	}
	//这样也行
	var a int = 66
	switch  {
	    case a == 66:
			fmt.Println("66")
	    case a<100:
		fmt.Println("55")
	    default:
		fmt.Println("default")
	}

	var b string = "wqil-jjio-m北京"
	//go中不同的遍历
	for i , v:= range b {
		//go中字符都是码表对应的数字组成的
		fmt.Println(i,"--->",v)
	}
	for i, v := range b {
		//以字符形式打印
		fmt.Printf("%d 是 %c\n" ,i,v)
	}
	j := 0
	sum :=0
	for i := 1  ; i <= 100; i++{
		if i % 9 == 0 {
			j++
			sum = sum + i
		}
	}
	fmt.Println(j,sum)
	f := 6
	for i := 0; i < 7; i++ {
		fmt.Println(i,"+",f,"=",i+f)
		f--
	}

}
