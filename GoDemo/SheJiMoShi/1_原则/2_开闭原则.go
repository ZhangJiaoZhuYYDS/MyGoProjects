package main

import "fmt"

// 抽象的业务员（执行某些业务的动作的集合），面向接口
type AbstractBanker interface {
	DoBusi() // 做业务
}

// 存款业务员
type SaveBanker struct {
	age int
}

// 实现业务接口
func (b *SaveBanker) DoBusi(i int) {
	b.age = i
	fmt.Println("进行了存款", b.age)
}

//转账业务员
type TranferBanker struct{}

/*类的改动是通过增加代码进行的，而不是修改源代码。*/

// 实现业务接口
func (b TranferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

///////////////////////////////////////////////////////////

// 对业务方法进行封装，实现传递不同类型，调用不同类型方法
//  把接口作为函数的参数, 可以动态的替换为接口的实现类,从而实现多态
func DoBusi(d AbstractBanker) {
	d.DoBusi()
}

func main() {
	// 实际的存款业务员
	saveBanker := SaveBanker{age: 10}
	// 调用同一业务方法，实现存款业务
	saveBanker.DoBusi(11)
	fmt.Println(saveBanker.age)
	// 实际的转账业务员
	tranferBanker := TranferBanker{}
	// 调用同一业务方法，实现转账业务
	tranferBanker.DoBusi()

	//////////////////////////////////////////////////////////
	// 注意，因为结构体实现接口方法是指针方法，所以传参时只能传指针类型的结构体；若结构体实现接口方法是指针方法，传参时可以传普通结构体，也可以传指针类型的结构体
	//DoBusi(&SaveBanker{})
	DoBusi(&TranferBanker{})
}
