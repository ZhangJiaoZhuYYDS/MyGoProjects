package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) test() {
	fmt.Println(a.Age)
}
func main() {
	var v *int
	fmt.Println(&v, v)

	animal := Animal{
		Name: "猫",
		Age:  6,
	}
	println(Animal{Name: "cat", Age: 5}.Age)
	a := animal
	b := &animal
	animal.Age = 0
	animal.test()
	a.test()
	b.test()
	println(">>>>>>>>>>>>>")
	println(Test(&i, &j))

}
var i int
var j int = 6
func Test(i *int, j *int) (int , int ) {
	*i++
	if *i == 5 {
		return *i,*j
	}
	Test(i,j)
	*j--
	return *i , *j
}