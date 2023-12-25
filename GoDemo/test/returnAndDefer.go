package main

import "fmt"

func MyDefer()  {
	fmt.Println("defer函数执行")
}
func MyReturn() int  {
	fmt.Println("return函数执行")
	return 666
}

func method1() int{
	defer MyDefer()
	return MyReturn()
}








func re() (i int){
	i = 0
	defer func( i int) {
		i++
		fmt.Println("defer函数执行，打印出》》》》",i)
	}(i)
	return func() int {
		fmt.Println("return函数先执行，打印出>>>>",123)
		return 66
	}()
}



func main() {
	//println(re())
	println(method1())
}
