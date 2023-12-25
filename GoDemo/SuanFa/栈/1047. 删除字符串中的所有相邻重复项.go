package main

func main() {
	removeDuplicates("abbaca")
}
func removeDuplicates(s string) string {
	b := []byte(s)
	stack := make([]byte, 0)
	// 添加第一个元素到栈顶
	stack = append(stack, s[0])
	for i := 1; i < len(b); i++ {
		if b[i] != b[i-1] {
			stack = append(stack, b[i])
		} else {
			stack = stack[0 : i-1]
			i++
		}
	}
	return string(b)
}
