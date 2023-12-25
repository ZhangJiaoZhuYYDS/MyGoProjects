package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name  string
	Age   int
	Money float64
	Hobby []string
}

// 实现sort接口

type IntSlice []int

func (x IntSlice) Len() int {
	return len(x)
}
func (x IntSlice) Less(i, j int) bool {
	return x[i] > x[j] //降序
}

// 交换位置
func (x IntSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	// 实现sort接口排序
	sl := IntSlice([]int{1, 5, 2, 9, 7})
	fmt.Println("排序前", sl)
	// sort按照Less定义规则排序
	sort.Sort(sl)
	fmt.Println("排序后", sl)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	p1 := Person{
		Name:  "li",
		Age:   10,
		Money: 100.1,
		Hobby: []string{"画画", "吃饭"},
	}
	p2 := Person{
		Name:  "li",
		Age:   9,
		Money: 100.2,
		Hobby: []string{"画", "吃", "玩"},
	}
	// 一个Person类型的切片
	persons := []Person{p1, p2}

	// 直接使用内置函数排序

	// 1 1、sort.Ints(x []int) sort.Float64s()
	// 切片元素的排序
	ints := []int{1, 4, 3, 2}
	fmt.Printf("升序%v\n", ints)
	sort.Ints(ints) // 升序
	fmt.Printf("%v\n", ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	// 2 sort.Slice(x any, less func(i, j int) bool) 切片元素的自定义排序
	// Slice函数有个好处，如果传入对象是切片，实现回调函数即可，如果传入对象是结构体，也可以自定义排序规则。
	sort.Slice(persons, func(i, j int) bool {
		//名字相同的情况下 按年龄逆序
		if persons[i].Name == persons[j].Name {
			return persons[i].Age > persons[j].Age
		}
		// 默认按年龄升序
		return persons[i].Name < persons[j].Name
	})

	// 3 sort.SearchInts(a []int , x int)int  二分查找
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	idx := sort.SearchInts(arr, 4)
	fmt.Printf("二分查找%v\n", idx) // 3

}
