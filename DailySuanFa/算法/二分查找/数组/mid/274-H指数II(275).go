// @Author zhangjiaozhu 2023/12/19 16:17:00
package mid

func hIndex2(citations []int) int {
	l := 1
	r := len(citations)
	ans := 0
	for l <= r {
		mid := l + (r-l)>>1

		if citations[len(citations)-mid] >= mid {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
