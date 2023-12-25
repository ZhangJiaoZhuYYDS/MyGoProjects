package main

// 344----反转字符串1  go函数库中没有反转切片的api
import "fmt"

func main() {
	s := "hello"
	bytes := []byte(s)
	reverseString(bytes)
	for i := range bytes {
		fmt.Printf("%c", bytes[i])
	}
}
func reverseString(s []byte) {

	if s == nil {
		return
	}
	// 左右指针不断往中间移动，同时交换两位置对应元素
	//初始化左右指针位置
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

}
