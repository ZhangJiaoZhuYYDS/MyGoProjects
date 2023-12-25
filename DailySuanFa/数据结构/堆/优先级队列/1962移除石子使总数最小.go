// @Author zhangjiaozhu 2023/12/23 10:54:00
package 优先级队列

import (
	"container/heap"
	"sort"
)

// 使用标准库
type PriorityQueue struct {
	sort.IntSlice
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.IntSlice[i] > pq.IntSlice[j]
}

func (pq *PriorityQueue) Push(v interface{}) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	arr := pq.IntSlice
	v := arr[len(arr)-1]
	pq.IntSlice = arr[:len(arr)-1]
	return v
}

func minStoneSum(piles []int, k int) int {
	pq := &PriorityQueue{piles}
	heap.Init(pq)
	for i := 0; i < k; i++ {
		pile := heap.Pop(pq).(int)
		pile -= pile / 2
		heap.Push(pq, pile)
	}
	sum := 0
	for len(pq.IntSlice) > 0 {
		sum += heap.Pop(pq).(int)
	}
	return sum
}

// 2 自定义实现堆
func minStoneSum2(piles []int, k int) int {
	n := len(piles)
	if n == 0 {
		return 0
	}
	sum := 0
	for _, p := range piles {
		sum = sum + p
	}
	for i := n / 2; i >= 0; i-- {
		adjustHeap(piles, i)
	}
	for i := 0; i < k; i++ {
		sum = sum - piles[0]/2
		piles[0] = piles[0] - piles[0]/2
		adjustHeap(piles, 0)
	}
	return sum
}

// 调整堆
func adjustHeap(piles []int, pos int) {
	val := piles[pos]
	for pos*2+1 < len(piles) {
		left := pos*2 + 1
		if left+1 < len(piles) && piles[left+1] > piles[left] {
			left++
		}
		if val < piles[left] {
			piles[pos] = piles[left]
			pos = left
		} else {
			break
		}
	}
	piles[pos] = val
}
