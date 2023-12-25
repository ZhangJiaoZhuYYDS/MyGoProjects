package main

import "fmt"

// 抽象层
type Car interface {
	RunCar()
}
type Driver interface {
	Drive(car Car)
}

// 实现层
type Benz struct {
}

func (b *Benz) RunCar() {
	fmt.Println("Banz is running")
}

type Zhang struct {
}

func (d *Zhang) Drive(car Car) {
	fmt.Println("张三 开车")
	car.RunCar()
}

// 业务层
func main() {
	var benz Car
	benz = new(Benz)
	var zhang3 Driver
	zhang3 = new(Zhang)
	zhang3.Drive(benz)

}
