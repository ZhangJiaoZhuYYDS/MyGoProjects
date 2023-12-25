package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip string
	Port int
}

//创建一个server接口
func NewServer(ip string, port int) *Server {
	Server := &Server{
		Ip: ip,
        Port: port,
    }
	return Server
	                                        }
func (this *Server) Handler(conn net.Conn) {
	// ...当前连接的业务
	fmt.Println("链接建立成功")
}
func (this *Server) Start()  {
	// socket listen
	listener , err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listener err",err)
        return
	}
	// close socket
	defer listener.Close()
    fmt.Println("server start")
    for {
       //accept
		conn, err := listener.Accept()
		if err!= nil {
        fmt.Println("listener.Accept err",err)
            continue
        }
		go this.Handler(conn)
	}
}