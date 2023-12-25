package main

import "fmt"

func main()  {
	/*1----------切片的创建*/
	fmt.Println(66666)
	a := [5]int{1,2,3,4,5}  //原始数组
	b := a[1:3]  // 1-----------创建数组切片
	c := []int{1,2,3}  // 另一种创建
	fmt.Println("原始数组",a)
	fmt.Println("切片",b)
	fmt.Println("伊利一种切片创建",c)

	/*2--------切片数据的修改  切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。*/
	for i , v := range b{
		fmt.Println(i,"-----",v)
		b[i]++
	}
	fmt.Println(a)
	/*3-------切片的长度和容量
	切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
	*/
	fmt.Println("切片的长度是",len(b),"切片的容量是",cap(b))
	fmt.Println(&a)
	o := 66
	fmt.Printf("变量的地址是 %x\n" , &o)
}