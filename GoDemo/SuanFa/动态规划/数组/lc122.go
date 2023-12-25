package main

import "fmt"

func main() {
	//prices := []int{7,1,5,3,6,4}
	prices2 := []int{7,6,5,4,3}
	maxProfit(prices2)
}
func maxProfit2(prices []int) int {
	// 计算所有增区间的和
	result := 0
	for i:=1 ; i < len(prices) ; i++ {
		if prices[i] - prices[i-1] > 0 {
			result+= prices[i] - prices[i-1]
		}
	}
	return result
}
func maxProfit(prices []int) int {
	// 动态规划
	if prices == nil || len(prices) < 2 {
		return 0
	}
	length := len(prices)
	dp := make([][2]int,length)

	// 初始化第一天持有股票和不持有股票的最大利润
	dp[0][0] = 0
	dp[0][1] =-prices[0]
	for i := 1 ; i < length ; i ++ {
		dp[i][0] = Max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1],dp[i-1][0]-prices[i])
	}
	for i := range dp {
		fmt.Println(">>>>>>>>>>>>>")
		for i2 := range dp[i] {
			fmt.Println(dp[i][i2])
		}
	}
	return dp[length-1][0]
}
func Max(a,b int) (int) {
	if a > b {
		return a
	}
	return b
}