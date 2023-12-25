package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quite := make(chan int)
	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Println("go________", i)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quite <- 0
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}
	ints := <-quite
	fmt.Println(ints)
}
