package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 声明函数类型变量
	var fun func()
	// 将匿名函数赋值给函数类型变量
	fun = func() {
		fmt.Println("匿名函数")
	}
	// 调用函数
	fun()

	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
		time.Sleep(1 * time.Second)
	}
	fmt.Println(666)
	fmt.Println(666)
	fmt.Println(666)
	wg.Wait()
}

/*输出结果：
5
5
5
5
5*/
