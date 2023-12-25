package main

/*切片，实际的是获取数组的某一部分，len切片<=cap切片<=len数组，切片由三部分组成：指向底层数组的指针、len、cap。*/
import "fmt"

func main() {
	//  1 声明一个切片，并且初始化，默认值是1 2 3 ，长度是 3
	// slice1 := []int{1,2,3}
	//  2 声明一个切片，但没有分配空间，通过make()赋值,默认值是0
	// var slice2 []int
	// slice2 = make([]int,3)
	// 3 声明切片，同时赋值
	// var slice3 = make([]int , 3)
	// 4
	slice4 := make([]int , 3)
	for index, value := range slice4 {
		fmt.Println(index,"---",value)
	}




	// 1 定义切片
	ints := make([]int, 8, 8)  //make声明一个切片,长度为8，容量为8
	var a []int   //直接声明
	// 2 切片初始化
	s := []int{4,5,6}
	// 初始化切片 s，是数组 arr 的引用。
	var arr = [3]int{4,5,6}
	a = arr[0:2]

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(ints)


	// 3 len()长度  cap() 容量

	// 4 空(nil)切片 一个切片在未初始化之前默认为 nil，长度为 0，

	// 5 切片截取（arr[0:2] 包前不包后）

	// 6 append()添加切片元素  copy()拷贝切片


}
