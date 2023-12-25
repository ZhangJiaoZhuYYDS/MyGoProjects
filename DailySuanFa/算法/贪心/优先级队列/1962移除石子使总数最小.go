// @Author zhangjiaozhu 2023/12/23 10:54:00
package 优先级队列

import (
	"container/heap"
	"sort"
)

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
