package main

import "fmt"

/*定义结构体*/
type Books struct {
	title string
	name string
	pages int
}
func main() {
	//创建一个新结构体
	book := Books{"文学", "javeweb", 100}
	fmt.Println(book)
	//访问结构体成员
	fmt.Println(book.name)
	//结构体作为函数参数
	printBook(book)
	//指针传参
	changeBook(&book)
	fmt.Println("指针传参修改值",book)

	//类方法(用指针类型)
	book.SetTitle("科幻")
	fmt.Println(book)
}
func printBook(book Books){
	fmt.Println(book.title)
	fmt.Println(book.name)
	fmt.Println(book.pages)
}
func changeBook(book *Books)  {
	book.title = "历史"
}

func (this *Books) SetTitle (newTitle string){
	this.title = newTitle
}
