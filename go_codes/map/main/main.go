package main

import "fmt"

func main() {
	// map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
	//1 map的创建方式一  make()
	personSalary := make(map[string]int)
	personSalary["tom"] = 200
	personSalary["eve"] = 100
	fmt.Println(personSalary["tom"])
	fmt.Printf("map的类型是 %T",personSalary)  //map的类型是 map[string]int
	//map的创建方式二  声明后必须初始化，不然

	//2  获取一个不存在的元素，map 会返回该元素类型的零值。
	//3  map 中某个 key 是否存在
	value,ok := personSalary["null"]  // value==0  ok==true
	fmt.Println(value,ok)
	//4 遍历map(有一点很重要，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。)
	for key, value := range personSalary {
		fmt.Println(key,"====",value)
	}
	//5 map之间不能用==判断，==只能用来检查map是否为nil;判断两个map是否相等是遍历比较map中的每个元素是否相等
	// 6 删除元素(没有返回值)
	delete(personSalary,"tom")
	fmt.Println(personSalary)



}
