package main

import "fmt"

// const 用来定义枚举类型
// iota 值自增
const (
	beijing = iota
	shanghai
	jiangsu
	nanjing
)
func main() {
	fmt.Println(beijing, shanghai,jiangsu)

}
