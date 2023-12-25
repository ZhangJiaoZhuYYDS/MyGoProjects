package main

import (
	"fmt"
	"time"
)

// 定义一个延时函数
func work(id int,done chan bool)  {
	fmt.Printf("协程%d开始运行\n",id)
	time.Sleep(time.Second)// 模拟耗时一秒
	fmt.Printf("协程%d结束运行\n",id)
	done <- true    // 协程结束利用通道发送信号给主协程
}

func main() {
	// 声明一个无缓存的通道
	done := make(chan bool)
	fmt.Println("主协程开始")
	// 开启三个协程
	for i := 0 ; i < 3 ; i++ {
		go work(i,done)
	}
	<- done  // 主协程阻塞等待子协程发送信号
	fmt.Println("主协程结束")
}
