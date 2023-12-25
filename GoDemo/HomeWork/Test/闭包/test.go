package main

import "fmt"

// 闭包修改共享的变量
func main() {
	n := 0
	f := func() int {
		n += 1
		return n
	}
	fmt.Println(f()) // 别忘记括号，不加括号相当于地址
	fmt.Println(f())
	fmt.Println(n)
}

/*
输出：
1
2
*/
