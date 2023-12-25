package main

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	println(needle[0:])
	println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	// 定义左端点
	l := 0
	// 定义一个新字符串
	s := ""
	// 遍历字符串
	for r := 0; r < len(haystack); r++ {
		if len(s) == len(needle) {
			if s == needle {
				return l
			}
			s = s[1:]
			l++
		}
		s += string(haystack[r])
		if s == needle {
			return l
		}
	}
	return -1
}
