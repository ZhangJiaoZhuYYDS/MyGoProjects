package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type World struct {
	
}

func (this *World)Hello(req string, res *string)  error {
	fmt.Println(req)
	*res = req + "server"
	return nil
}

func main() {
	// 1 注册rpc
	err := rpc.RegisterName("RPC",new(World))
	if err != nil {
		fmt.Println("注册rpc失败"+err.Error())
		return
	}
	// 2 设置监听
	con, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net listen err:"+err.Error())
		return
	}
	defer con.Close()

	fmt.Println("开始监听")

	// 3 建立链接
	for{
		//接收连接
		accept, err := con.Accept()
		if err != nil {
			fmt.Println("accept err:"+err.Error())
			return
		}
		// 4 绑定服务
		go rpc.ServeConn(accept)
	}
}
