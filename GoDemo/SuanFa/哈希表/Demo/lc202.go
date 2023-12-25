package main

// 202 快乐数
func main() {
	//n := 19
	println(isHappy(2))
	println(getSum(2))
}

func isHappy(n int) bool {
	m := map[int]bool{}
	// 判断n是否等于1且m不含有该元素,然后循环调用getSum函数
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}
func getSum(n int) int {
	sum := 0 //记录每次计算后数的结果
	for n > 0 {
		// 记录余数的平方
		sum += (n % 10) * (n % 10)
		// n等于n的10位数
		n = n / 10
	}
	return sum
}
