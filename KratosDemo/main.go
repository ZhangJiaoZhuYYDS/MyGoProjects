package main

import "fmt"

type Message string

type Channel struct {
	Message Message
}

type BroadCast struct {
	Channel Channel
}

func NewMessage() Message {
	return Message("hello")
}
func NewChannel(m Message) Channel {
	return Channel{Message: m}
}

func NewBroadCast(c Channel) BroadCast {
	return BroadCast{
		Channel: c,
	}
}
func (c Channel) GetMsg() Message {
	return c.Message
}
func (b BroadCast) start() {
	msg := b.Channel.Message
	fmt.Println(msg)
}

func main() {
	//Message := NewMessage()
	//channel := NewChannel(Message)
	//broadCast := NewBroadCast(channel)
	//broadCast.start()
	
}
