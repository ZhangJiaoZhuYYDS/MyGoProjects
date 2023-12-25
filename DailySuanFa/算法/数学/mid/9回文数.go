// @Author zhangjiaozhu 2023/12/23 10:15:00
package mid

// 同 7

func isPalindrome(x int) bool {
	source, result := x, 0
	for x > 0 {
		n := x % 10
		result = result*10 + n
		x /= 10
	}

	return source == result
}
