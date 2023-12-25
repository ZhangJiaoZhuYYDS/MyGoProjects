package main

import (
	"container/heap"
	"sort"
)

type PriorityQueue struct {
	sort.IntSlice
}

// 实现排序包接口，把数组修改为从大到小的排列顺序，目的实现大顶堆
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.IntSlice[i] > pq.IntSlice[j]
}

// 堆的入队列
func (pq *PriorityQueue) Push(v interface{}) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

// 堆的出队列
func (pq *PriorityQueue) Pop() interface{} {
	arr := pq.IntSlice
	v := arr[len(arr)-1]
	pq.IntSlice = arr[:len(arr)-1]
	return v
}

func minStoneSum(piles []int, k int) int {
	pq := &PriorityQueue{piles}
	//大顶堆的初始化操作
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

func main() {
	piles := []int{4, 3, 6, 7}
	k := 3
	minStoneSum2(piles, k)

}
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
