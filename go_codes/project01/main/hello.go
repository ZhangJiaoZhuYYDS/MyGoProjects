package main

/*总结：
      go build xxx.go 生成一个exe文件
	  go build -o myGo.exe xxx.go 生成指定文件名的exe文件
      go run 只编译运行

		Go程序开发注意事项（重点）
         1)G0源文件以"go”为扩展名。
         2)Go应用程序的执行入口是main()函数。
         3)G0语言严格区分大小写。
         4)G0方法由一条条语句构成，每个语句后不需要分号(G0语言会在每行后自动加分号)，这也体现出Golang的简洁性。
         5)G0编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一个，否则报错
         6)go语言定义的变量或者import的包如果没有使用到，代码不能编译通过。
         7)大括号都是成对出现的，缺一不可。



main 包的初始化顺序为：

首先初始化被导入的包。因此，首先初始化了 rectangle 包。
接着初始化了包级别的变量 rectLen 和 rectWidth。
调用 init 函数。
最后调用 main 函数。
*/
import (
	"fmt"
	"go_codes/project01/main/new"  //引入自定义包文件（包文件内方法必须大写）
	_ "log"   //空白标识符  引入包但暂时不使用它
)

var d int = 66

func init()  {
	if d > 65 {
		fmt.Println("初始化主函数内init（）函数hello init")
	}
}
func number() int {
	return d*2
}
func main() {
	fmt.Println("hello world!\r张飞")
	fmt.Println("太牛龙\r张飞")
	a , b := 21,23
	b , c := 21,21
	fmt.Println(a,b,c)
	new.New()

	switch num := number(); {
	case num<5:
		fmt.Println("num 小于5")
		fallthrough
	case num<200:
		fmt.Println("num小于200")
		fallthrough
	case num<500:
		fmt.Println("num 小于500")
	}


}
