package main

import (
	"sync"
)

var c = make(chan int)
var g sync.WaitGroup

func main() {
	num := 0
	g.Add(1)
	go add(c)
	for {
		select {
		case i, ok := <-c:
			if !ok {
				g.Done()
				return
			}
			num += i
		}
	}

}
func add(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
