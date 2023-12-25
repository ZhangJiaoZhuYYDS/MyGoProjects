package main

import "fmt"

func main() {
	mapTable := make(map[int]int)
	mapTable[0]=1;
	mapTable[0]=1;
	if i,ok := mapTable[0];ok{
		fmt.Println(i,ok)
	}
	for i, i2 := range mapTable {
		fmt.Println(i,i2)
	}

	fmt.Println(len(mapTable))

}
