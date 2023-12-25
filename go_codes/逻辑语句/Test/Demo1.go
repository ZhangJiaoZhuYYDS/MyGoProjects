package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//for i:=0 ;  i < 3 ; i++ {
	//	var sum int = 0
	//	for j := 0; j < 3; j++ {
	//		var score int = 0
	//		fmt.Printf("请输入第%d个班,第%d个同学的成绩",i ,j)
	//		fmt.Scanln(&score)
	//		sum+=score
	//	}
	//	fmt.Printf("第%d个班的平均成绩是%d",i,sum/3)
	//}

	//我们使用随机数，还需要给rand设置一个种子
	//time.Now().Unix():返回一个从1970：01：01的0时0分o秒到现在的纳秒数
	rand.Seed(time.Now().UnixNano())

	c := rand.Intn(100)+1 // [0 100)
	fmt.Println(c)

	//goto的使用
	fmt.Println(1)
	goto label1
	fmt.Println(2)
	fmt.Println(3)
	label1:
	fmt.Println(4)
	fmt.Println(5)

}
