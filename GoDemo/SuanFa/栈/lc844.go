package main

// 844 比较含退格的字符串  切片模拟栈
//    s :=  []string
//    v 入栈：append(s,v) 添加到尾部(这里只是添加到了尾部，并没有像正常栈压入头部)
//    v 出栈：s = s[:len(s)]
//    记得len(s)==0的时候就不能再出栈了
func main() {
	s := "ab#c"
	t := "ad#c"
	println(backspaceCompare(s, t))

}

// 方法一 字符串/双指针

// 方法二 模拟栈
func backspaceCompare(s string, t string) bool {
	return test(s) == test(t)
}
func test(str string) string {
	// 声明一个空字节数组(模拟栈)
	s := []byte{}
	// 遍历字符串
	for i := range str {
		v := str[i]
		// 字符不是‘#’，就把当前字符添加到切片s
		if v != '#' {
			s = append(s, v)
			// 切片不为空的情况下，截取切片
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}
