package main

import "fmt"

type Car interface {
	Run()
}
type Driver interface {
	Drive( car Car)
}

type Benz struct {

}

func (Benz)Run()  {
	fmt.Println("benz开动了")
}

type Zhang struct {

}

func (z Zhang)Drive(car Car)  {
	car.Run()
}
func test()  {
	
}

func main() {
	Zhang{}.Drive(Benz{})
}
