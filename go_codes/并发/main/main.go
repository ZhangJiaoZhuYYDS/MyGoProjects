package main

import (
	"fmt"
	"time"
)

//func say(s string)  {
//	for i := 0; i < 5; i++ {
//		time.Sleep(1*time.Millisecond) //每次循环让线程休眠100微秒
//		fmt.Println(s)
//	}
//
//}



func main() {
	//go say("go线程")
	//say("主线程")
	go func() {
		fmt.Println("6666")
	}()
	func(){
		time.Sleep(1*time.Second)
	}()
}
