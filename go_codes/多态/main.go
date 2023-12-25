package main



//interface接口和interface{}万能数据类型以及断言判断数据类型
import (
	"fmt"


)

//声明接口
type Animal interface {
	Sleep()
	GetColor() string
}
//猫类
type Cat struct {
	color string
}
//猫类实现接口重写方法
func (this *Cat) Sleep()  {
	fmt.Println("cat sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
//狗类
type Dog struct {
	color string
}
//狗类实现接口重写方法
func (this *Dog) Sleep()  {
	fmt.Println("dog sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}

func main() {
	var animal Animal
	//多态：调用猫类方法
	animal = &Cat{"Green"}
	animal.Sleep()
	fmt.Println("猫颜色",animal.GetColor())
	//多态：调用狗类方法
	animal = &Dog{"yellow"}
	animal.Sleep()
	fmt.Println("狗颜色",animal.GetColor())

	Test("666")

}
func Test(arg interface{})  {
	value , ok := arg.(string)
	if ok {
		fmt.Println(value,ok,"string")
	}else {
		fmt.Println(value,ok,"string")
	}
}
