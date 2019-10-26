package pacman

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
	"image"
)

type player struct {
	images 			[8]*ebiten.Image
	currentImage 	int
	currPos 		pos
}


func newPlayer(x,y int) *player {
	p := &player{}
	p.currPos = pos{x,y}
	p.currentImage = 0
	p.loadImages()
	return p
}

func (p *player) loadImages() {
	for i:=0; i<8; i++ {
		img, _, err := image.Decode(bytes.NewReader(pacimages.PlayerImages[i]))
		handleError(err)
		p.images[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
}

func (p *player) getCurrentImage() *ebiten.Image{
	return p.images[p.currentImage]
}

func (p *player) draw(sc *ebiten.Image){
	v := p.currPos
	x := float64(v.y * stageBlocSize)
	y := float64(v.x * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x,y)
	sc.DrawImage(p.getCurrentImage(), op)
}







