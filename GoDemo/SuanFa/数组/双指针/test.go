package main

import "fmt"

func main() {
	var s *int
	s = new(int)
	*s = 9
	fmt.Println(*s)
	arr := make([]int, 3, 9)
	arr1 := new([]int)
	fmt.Printf("%T,%T", s, arr)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr1)
	fmt.Println(arr1 == nil)
}
