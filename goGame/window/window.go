package window

import (
	"MyGoProject/goGame/config"
	"MyGoProject/goGame/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct {
	input *input.Input
	conf  *config.Config
}

func NewGame() *Game {
	return &Game{
		input: &input.Input{Msg: "hello"},
		conf: config.LoadConfig(),
	}
}

func (g *Game)	Update() error  {
	g.input.Update()
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen,g.input.Msg)
	screen.Fill(g.conf.BgColor)
}
func (g *Game) Layout (outsideWidth , outsideHeight int) (screenWidth,screenHeight int) {
	return g.conf.ScreenWidth / 2,g.conf.ScreenHeight / 2
}

func init() {
	game := NewGame()
	cfg := config.LoadConfig()
	// 1 初始化
	ebiten.SetWindowSize(cfg.ScreenWidth,cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	if err := ebiten.RunGame(game); err !=nil {
		log.Fatal(err)
	}

}
