package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓存通道
	controlChan := make(chan struct{})

	// 启动一个协程
	go func() {
		for {
			select {
			case <-controlChan:
				// 收到通知，执行相应操作
				fmt.Println("协程收到通知，开始执行操作...")
				time.Sleep(1 * time.Second)
				fmt.Println("协程操作完成，等待下一次通知...")
			}
		}
	}()

	// 在主协程中，向无缓存通道发送信号，以控制另一个协程
	for i := 0; i < 3; i++ {
		fmt.Println("主协程发送通知...")
		controlChan <- struct{}{}
		time.Sleep(2 * time.Second)
	}

	// 为了便于观察效果，让主协程等待一会儿
	time.Sleep(5 * time.Second)
}
/*首先创建了一个无缓存通道 controlChan，然后启动了一个新的协程。这个协程会一直等待接收来自 controlChan 的信号。
当主协程向 controlChan 发送信号时，另一个协程会收到通知并执行相应的操作。通过这种方式，无缓存通道可以实现协程之间的同步和控制。*/