package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for i2 := range os.Args {
			log.Println(i2)
		}
	}
	print(55)
	defer print(12)
	file, err := ioutil.ReadFile("./test03.go")
	if err != nil {
		return
	}
	print(string(file))
	print(123)
	from := io.ReaderFrom(os.Stdin)
	fmt.Println(from)

}
