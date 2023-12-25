// @Author zhangjiaozhu 2023/12/21 11:42:00
package 前缀和

func maximumSumOfHeights(maxHeights []int) int64 {
	// 统计长度
	n := len(maxHeights)
	// 分别是左右最大前缀和 fl[i] 表示从左起到第i个元素组成的最大和  fr 表示从右起到第i个元素组成的最大和
	fl, fr := make([]int64, n), make([]int64, n)
	// 分别是前缀和的单调栈（存的下标）
	st1, st2 := []int{}, []int{}

	//左起计算前缀和
	for i := 0; i < n; i++ {
		// 栈不空，就判断当前下索引元素和maxHeights中的单调栈索引对应的元素的大小，小的话就栈顶出栈
		for len(st1) > 0 && maxHeights[i] < maxHeights[st1[len(st1)-1]] {
			st1 = st1[:len(st1)-1]
		}
		//栈为空，就说明当前索引元素比左边的元素小，就保存左起前i的前缀和为（i+1）*当前元素   通俗说就是：以当前索引对应元素为峰顶，发现它左边的元素比他大，就把左边的元素统一削减为峰顶元素大小一样
		if len(st1) == 0 {
			fl[i] = int64(i+1) * int64(maxHeights[i])
			// 不为空，就说明当前索引元素比左边的元素大，直接取出栈顶索引找到前缀和+上当前索引作为左起前i的前缀和
		} else {
			last := st1[len(st1)-1]
			fl[i] = fl[last] + int64(i-last)*int64(maxHeights[i])
		}
		// 当前索引入栈
		st1 = append(st1, i)
	}
	//右起计算前缀和
	for i := n - 1; i >= 0; i-- {
		for len(st2) > 0 && maxHeights[i] < maxHeights[st2[len(st2)-1]] {
			st2 = st2[:len(st2)-1]
		}
		if len(st2) == 0 {
			fr[i] = int64(n-i) * int64(maxHeights[i])
		} else {
			last := st2[len(st2)-1]
			fr[i] = fr[last] + int64(last-i)*int64(maxHeights[i])
		}
		st2 = append(st2, i)
	}
	ans := int64(0)
	// fmt.Printf("%d\n%d\n", fl, fr)
	// 比较左右前缀和寻找最大和
	for i := 0; i < n; i++ {
		ans = max(ans, fl[i]+fr[i]-int64(maxHeights[i]))
	}
	return int64(ans)
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
