package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster 广播器，它使用局部变量 clients 来记录当前连接的客户集合，每个客户唯一被记录的信息是其对外发送消息通道的 ID

type client chan<- string // 对外发送消息的通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string) // 所有连接的客户端
)

func broadcaster() {
	// 这里使用一个字典来保存用户 clients，字典的 key 是各连接申明的单向并发队列
	clients := make(map[client]bool)
	/*
		使用一个 select 开启一个多路复用：
		每当有广播消息从 messages 发送进来，都会循环 cliens 对里面的每个 channel 发消息。
		每当有消息从 entering 里面发送过来，就生成一个新的 key - value，相当于给 clients 里面增加一个新的 client。
		每当有消息从 leaving 里面发送过来，就删掉这个 key - value 对，并关闭对应的 channel。
	*/

	for {
		select {
		// 从消息通道中获取消息
		case msg := <-message:
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}
		// 接收进入通道中的通道，把他放进map集合，状态设为true
		case cli := <-entering:
			clients[cli] = true
		// 接收离开通道中的通道，把他从map删除
		case cli := <-leaving:
			delete(clients, cli)
		}

	}
}

// handleConn 函数会为每个过来处理的 conn 都创建一个新的 channel，开启一个新的 goroutine 去把发送给这个 channel 的消息写进 conn
func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道
	go clientWriter(conn, ch)

	// 获取连接过来的 ip 地址和端口号；
	who := conn.RemoteAddr().String()
	// 把欢迎信息写进 channel 返回给客户端
	ch <- "欢迎" + who
	// 生成一条广播消息写进 messages 里
	message <- who + "上线le"
	// 把这个 channel 加入到客户端集合，也就是 entering <- ch；
	entering <- ch
	// 监听客户端往 conn 里写的数据，每扫描到一条就将这条消息发送到广播 channel 中
	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + "<<<<<<<<" + input.Text()
	}
	// 如果关闭了客户端，那么把队列离开写入 leaving 交给广播函数去删除这个客户端并关闭这个客户端
	leaving <- ch
	// 广播通知其他客户端该客户端已关闭；
	message <- who + "下线"
	// 最后关闭这个客户端的连接 Conn.Close()。
	err := conn.Close()
	if err != nil {
		return
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for s := range ch {
		fmt.Fprintln(conn, s)
	}
}
