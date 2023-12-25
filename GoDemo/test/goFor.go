package main

import (
	"fmt"
	"time"
)

func main() {

	ints := make(chan int)
	go func(ints chan int) {
		for i := 0; i < 5; i++ {
			time.Sleep(1*time.Second)
			 ints <- i
		}
		close(ints)
	}(ints)

	for i := range ints {
		fmt.Println(i)
	}


}
