package pacman

import "github.com/hajimehoshi/ebiten"

type Game struct {
	scene *scene
}

type pos struct {
	x,y int
}

func NewGame() *Game{
	g := &Game{}
	g.scene = newScene(nil);
	return g
}

func (g *Game) ScreenWidth() int {
	return g.scene.ScreenWidth()
}

func (g *Game) ScreenHeight() int {
	return g.scene.ScreenHeight()
}

func (g *Game) Update(screen *ebiten.Image) error{
	return g.scene.update(screen)
}

