package main

// 541----反转字符串2
func main() {

}
func reverseStr(s string, k int) string {
	b := []byte(s) //将字符串转为字节切片

	// 遍历 ，每次移动2k个单位
	for i := 0; i < len(b); i += 2 * k {
		// 每隔2k个字符 对前k个字符进行反转
		// 剩余字符小于2k,但大于或等于k，就反转前k个字符
		if i+k < len(b) {
			reverseSlice(b[i : i+k])
			// 剩余字符小于k，就反转剩余所有字符
		} else {
			reverseSlice(b[i:len(b)])
		}
	}
	return string(b)
}
func reverseSlice(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
