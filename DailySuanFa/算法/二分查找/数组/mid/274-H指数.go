// @Author zhangjiaozhu 2023/12/19 16:16:00
package mid

import "sort"

func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return
}
