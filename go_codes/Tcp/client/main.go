package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "0.0.0.0:8888")
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sucess",dial)

}
