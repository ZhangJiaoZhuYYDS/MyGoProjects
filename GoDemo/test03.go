// @Author zhangjiaozhu 2023/3/26 19:38:00
package main

import "fmt"

type Node struct {
	value int
	next *Node
}
func traverseListRecursive(node *Node) []int {
	if node == nil {
		return []int{}
	}
	//array = append(array, node.value)
	//return traverseListRecursive(node.next, array)
	recursive := traverseListRecursive(node.next)
	recursive = append(recursive,node.value)
	return recursive
}

func main()  {
	head2 := Node{
		value: 2,
		next:  nil,
	}
	head1 := Node{
		value: 1,
		next:  &head2,
	}
	head0 := Node{
		value: 0,
		next:  &head1,
	}
	//arr := make([]int,0)
	recursive := traverseListRecursive(&head0)
	fmt.Println(recursive)
}

//func reversePrint(head *Node) []int {
//	var my func(node *Node , array []int) []int
//	my = func (node *Node, array []int) []int {
//		if node == nil {
//		return array
//	}
//		array = append(array, node.value)
//		return traverseListRecursive(node.next, array)
//	}
//
//}