package main

import "fmt"
//父类
type Man struct {
	Name string
}

//父类方法
func (this *Man) GetName()  {
	fmt.Println("父类get方法")
}

func (this *Man) SetName(name string)  {
	this.Name = name
}

type SuperMan struct {
	Man    //SuperMan 继承 Man
	age int
}
//子类重写父类方法
func (this *SuperMan) SetName( name string) {
	this.Name = name
}
func (this *SuperMan) GetAge()  {
	fmt.Println("子类方法")
}
func main() {
	m := Man{"父类"}
	m.GetName()
	m.SetName("新父类")
	fmt.Println(m)
	fmt.Println("................")
	//声明子类对象方法一
	//s := SuperMan{Man{"子类"},16}
	//声明子类方法二
	var s SuperMan
	s.Name = "子类"
	s.age = 13
	fmt.Println(s)
	s.GetName()
	s.SetName("新子类")
	fmt.Println(s)

}
