// @Author zhangjiaozhu 2023/7/28 9:53:00
package mid

import (
	"fmt"
	"sort"
)

// 当需要双层for循环遍历的时候，可以考虑是否需要遍历所有的情况，可以使用二分查找来快速找到匹配条件，避免重复机械

func findCountOfSuccessfulSpells(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))

	for i := 0; i < len(spells); i++ {
		// 左指针
		left := 0
		// 右指针
		right := len(potions) - 1
		// 记录前一位不满足条件的
		index := len(potions)

		for left <= right {
			mid := (right-left)/2 + left
			if int64(spells[i])*int64(potions[mid]) >= success {
				right = mid - 1
				index = mid
			} else {
				left = mid + 1
			}
		}

		res[i] = len(potions) - index
	}

	return res
}

func main() {
	spells := []int{2, 4, 6}
	potions := []int{1, 3, 5, 7, 9}
	success := int64(10)

	res := findCountOfSuccessfulSpells(spells, potions, success)
	fmt.Println(res)
}
