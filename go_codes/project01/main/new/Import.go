package new

import "fmt"

var baovar int = 64

func init()  {
	fmt.Println("先初始化引入包内中的init函数=====",baovar)

}
func New()  {
	fmt.Println("New")
}
