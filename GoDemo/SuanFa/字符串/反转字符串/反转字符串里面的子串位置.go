package main

// lc151
func main() {
	// 输入:
	s := "the sky is blue"
	// 输出: "blue is sky the"
	println(reverseWords(s))
}
func reverseWords(s string) string {
	// 新建切片result  []byte{}
	result := []byte{}
	// 将字符串转为字节切片
	b := []byte(s)
	// 获取切片长度
	lenB := len(b)
	// 定义双指针 l r 都从最后一位开始 l每次移动一位，r达到条件再移动
	l, r := lenB-1, lenB-1
	// 遍历切片元素 (  l >= 0 )
	for l >= 0 {
		// 如果l元素为‘ ’ ，就把s[l+1:r+1],s[l]添加到新切片 然后l-- ;r = l
		if b[l] == ' ' {
			// 添加两指针之间的元素
			result = append(result, b[l+1:r+1]...)
			// 添加空格
			result = append(result, b[l])
			l--
			r = l
		}
		// 如果l == 0 就把s[l:r+1]添加到result
		if l == 0 {
			result = append(result, b[l:r+1]...)
			l--
		}
	}
	return string(result)

}
