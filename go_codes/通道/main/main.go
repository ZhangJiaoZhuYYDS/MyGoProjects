package main

import "fmt"

func sum(a []int, c chan int)  {
	sum := 0
	for _, i2 := range a {
		sum+=i2
	}
	c <- sum   //3（写入信道）  把值发送给通道c
}
func main() {
	s := []int{1,2,3,4}   //初始化切片
	ints := make(chan int) // 1   make生成初始化通道
	go sum(s[:len(s)/2],ints)
	go sum(s[len(s)/2:],ints)
	x , y := <- ints , <-ints //2  读取通道中接收c
	fmt.Printf("通道中的sum1是%v,通道中的sum2是%v,sum1+sum2=%d",x,y,x+y)

	/* 4 发送与接收默认是阻塞的
	发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。*/

	// 5 死锁

}
