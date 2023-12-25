package main

import "fmt"

func main() {
	 arr :=[]int {73, 74, 75, 71, 71, 72, 76, 73}
	 fmt.Println(dailyTemperatures(arr))

}
func dailyTemperatures(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
