package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	//声明一个map
	HashTable := make(map[int]int)
	//遍历数组nums
	for i, v := range nums {
		//判断map是否target-v
		if value , ok := HashTable[target-v]; ok{
			//如果存在就返回两个数的下标组成的数组
			return []int{ value ,i}
		}
		// 不存在就放入map
		HashTable[v] = i
	}
	//遍历完也不存在返回nil
	return nil

}

func main() {
	sum := TwoSum([]int{2, 5, 6, 7}, 9)
	for _, i := range sum {
		fmt.Println(i)
	}
}
