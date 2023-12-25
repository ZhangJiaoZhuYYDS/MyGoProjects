package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start.........")
	// 1 tcp 使用tcp网络协议  2 监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err",err)
		return
	}
	//延时关闭监听
	defer listen.Close()

	//循环等待客服端链接
	for  {
		fmt.Println("等待客服端链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() error",err)
		}else {
			fmt.Printf("accept() sucess",conn)
		}
	}
}
