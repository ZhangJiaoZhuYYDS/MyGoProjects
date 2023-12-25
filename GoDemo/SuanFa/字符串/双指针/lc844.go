package main

//844 比较含退格的字符串
func main() {
	s := "###"
	t := "###"
	println(backspaceCompare(s, t))
}
func backspaceCompare(s string, t string) bool {
	// 倒序初始化两个字符串的指针位置
	l, r := len(s)-1, len(t)-1
	// 初始化两个字符串遇到‘#’的跳过数量
	skipS, skipT := 0, 0

	// 循环判断两指针是否都到达了0
	for l >= 0 || r >= 0 {
		// 判断s字符串是否到达0
		for l >= 0 {
			if s[l] == '#' {
				skipS++
				l--
			} else if skipS > 0 {
				skipS--
				l--
			} else {
				break
			}
		}
		// 判断t字符串是否到达0
		for r >= 0 {
			if t[r] == '#' {
				skipT++
				r--
			} else if skipT > 0 {
				skipT--
				r--
			} else {
				break
			}
		}
		// 每次等到两字符串skip清零后，能确定一个字符时，进行字符比较
		if l >= 0 && r >= 0 {
			if s[l] != t[r] {
				return false
			}
			// 防止都是‘#’
		} else if l >= 0 || r >= 0 {
			return false
		}
		l--
		r--
	}
	return true
}
