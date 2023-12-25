package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。*/

func main() {
	var s string = "hello world"
	fmt.Printf("%c",s[0])
	fmt.Println(strings.Split(s,","))

	// 1 遍历字符串中的字符（%c依字符方式打印）
	for i, i2 := range s {        // i 是 索引 ； i2 是 字符对应的字符码值
		fmt.Println(i,"---->",i2)
		fmt.Printf("%v--->%c\n" ,i,i2)   // %c 依字符方式打印
	}
	// 2 rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
	runes := []rune(s)
	for _, i2 := range runes {
		fmt.Println(i2)
	}
	//3 用字节切片构建字符串
	byteSlice := []byte{99,97,95,99}
	s2 := string(byteSlice)
	fmt.Println(s2)
	// 4 字符串的长度
	fmt.Println(utf8.RuneCountInString(s2))
	// 5 Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改。
}
