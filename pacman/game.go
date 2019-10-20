package pacman

import "github.com/hajimehoshi/ebiten"

type Game struct {
	scene *scene
}

func NewGame() *Game{
	g := &Game{}
	return g
}

func (g *Game) ScreenWidth() int {
	return 320
}

func (g *Game) ScreenHeight() int {
	return 240
}

func (g *Game) Update(screen *ebiten.Image) error{
	if(g.scene == nil){
		g.scene = newScene()
	}
	return g.scene.update(screen)
}

