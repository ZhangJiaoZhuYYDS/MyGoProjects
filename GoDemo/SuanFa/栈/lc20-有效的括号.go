package main

func main() {

}
func isValid(s string) bool {
	mp := map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := make([]byte, 0)
	if s == "" {
		return false
	}
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == mp[s[i]] {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
