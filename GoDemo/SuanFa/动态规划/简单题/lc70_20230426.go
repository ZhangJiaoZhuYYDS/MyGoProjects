package main
// 爬楼梯
func main() {

}
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	// 声明记录爬到第几层需要多少种方案的数组
	dp := make([]int,n+1)
	//初始化，爬到第一层需要一种方案
	dp[1] = 1
	// 爬到第二层需要2种方案
	dp[2] = 2
	// 计算爬到每一次爬到第n层的方案，添加到数组中
	for i := 3 ; i<= n ; i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}
/*java 滚动数组*/
/*class Solution {
    // 滚动数组，斐波那契
    public int climbStairs(int n) {
       int p = 0, q = 0 , r = 1;
       for(int i = 1 ; i <=n;i++){
           p = q;
           q = r;
           r = p+q;
       }
       return r;
    }
}*/