// @Author zhangjiaozhu 2023/9/22 10:03:00
package 排序

import (
	"github.com/Arafatk/DataViz/tree/master/trees/redblacktree"
	"math"
	"sort"
)

// 414  排序/有序数组/动态维护数组

/*
方法一：排序
1 将数组从大到小排序后，从头开始遍历数组，
2 通过判断相邻元素是否不同，来统计不同元素的个数。如果能找到三个不同的元素，就返回第三大的元素，否则返回最大的元素。
*/
func thirdMax(nums []int) int {
	// 排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // 把数组转换成sort.intslice类型，然后降序排序
	count := 1
	for i := range nums {
		// 从第二位开始
		if i >= 1 {
			if nums[i] != nums[i-1] {
				count += 1
				if count == 3 {
					return nums[i]
				}
			}
		}
	}
	return nums[0]
}

/*
方法二：有序集合
我们可以遍历数组，同时用一个有序集合来维护数组中前三大的数。
具体做法是每遍历一个数，就将其插入有序集合，若有序集合的大小超过 3，就删除集合中的最小元素。
这样可以保证有序集合的大小至多为 3，且遍历结束后，若有序集合的大小为 3，其最小值就是数组中第三大的数；若有序集合的大小不足 333，那么就返回有序集合中的最大值。
*/
func thirdMax2(nums []int) int {
	// 构建红黑树
	t := redblacktree.NewWithIntComparator()
	for _, num := range nums {
		t.Put(num, nil)
		if t.Size() > 3 {
			t.Remove(t.Left().Key)
		}
	}
	if t.Size() == 3 {
		return t.Left().Key.(int)
	}
	return t.Right().Key.(int)
}

// java
//class Solution {
//public int thirdMax(int[] nums) {
//TreeSet<Integer> s = new TreeSet<Integer>();
//for (int num : nums) {
//s.add(num);
//if (s.size() > 3) {
//s.remove(s.first());
//}
//}
//return s.size() == 3 ? s.first() : s.last();
//}
//}

/*
我们可以遍历数组，并用三个变量 aaa、bbb 和 ccc 来维护数组中的最大值、次大值和第三大值，以模拟方法二中的插入和删除操作。为方便编程实现，我们将其均初始化为小于数组最小值的元素，视作「无穷小」，比如 −263-2^{63}−2
63

	等。

遍历数组，对于数组中的元素 num\textit{num}num：

若 num>a\textit{num}>anum>a，我们将 ccc 替换为 bbb，bbb 替换为 aaa，aaa 替换为 num\textit{num}num，这模拟了将 num\textit{num}num 插入有序集合，并删除有序集合中的最小值的过程；
若 a>num>ba>\textit{num}>ba>num>b，类似地，我们将 ccc 替换为 bbb，bbb 替换为 num\textit{num}num，aaa 保持不变；
若 b>num>cb>\textit{num}>cb>num>c，类似地，我们将 ccc 替换为 num\textit{num}num，aaa 和 bbb 保持不变；
其余情况不做处理。
遍历结束后，若 ccc 仍然为 −263-2^{63}−2
63

	，则说明数组中不存在三个或三个以上的不同元素，即第三大的数不存在，返回 aaa，否则返回 ccc。
*/
func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v > a {
			c = b
			b = a
			a = v
		} else if a > v && v > b {
			c = b
			b = v
		} else if b > v && v > c {
			c = v
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
