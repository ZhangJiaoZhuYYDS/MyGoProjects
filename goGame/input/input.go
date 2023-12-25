package input
/*下面我们来监听键盘的输入，当前只处理3个键：左方向←，右方向→和空格。

ebiten提供函数IsKeyPressed来判断某个键是否按下，同时内置了100多个键的常量定义，见源码keys.go文件。ebiten.KeyLeft表示左方向键，ebiten.KeyRight表示右方向键，ebiten.KeySpace表示空格。
————————————————
*/
import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	Msg string
}

func (input *Input) Update(){
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<")
		input.Msg ="left pressed"
	}else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>")
		input.Msg ="right pressed"
	}else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println(".......................")
		input.Msg ="space pressed"
	}
}

