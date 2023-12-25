package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	con, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	// 2 客户端退出关闭连接
	defer con.Close()

	// 3 调用远程服务器函数
	var reply string
	err = con.Call("RPC.Hello", "客户端", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}
