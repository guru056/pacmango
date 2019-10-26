package pacman

import (
	"github.com/hajimehoshi/ebiten"
)

//define struct ghost
//define constructor
//define getCurrentImage func
//define draw func


type ghost struct {
	kind		elem
	currPos		pos
	currentImg	int
}

func newGhost(x,y int, kind elem) *ghost {
	g 			:= &ghost{}
	g.kind 		= kind
	g.currPos 	= pos{x,y}
	return g
}

func (g *ghost) getCurrentImage(imgs []*ebiten.Image) *ebiten.Image {
	return imgs[g.currentImg]
}

func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image ) {
	v := g.currPos
	x := float64(v.y * stageBlocSize)
	y := float64(v.x * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x,y)
	screen.DrawImage(g.getCurrentImage(imgs), op)
}