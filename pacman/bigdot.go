package pacman

import (
	"bytes"
	"container/list"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)
type bigDotManager struct {
	dots *list.List
	images [2]*ebiten.Image
}

func newBigDotManager() *bigDotManager {
	bd := &bigDotManager{}
	bd.dots = list.New()
	bd.loadImages()
	return bd
}

func (bd *bigDotManager) loadImages() {
	img, _, err := image.Decode(bytes.NewReader(pacimages.BigDot1_png))
	handleError(err)
	bd.images[0], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handleError(err)

	img2, _, err := image.Decode(bytes.NewReader(pacimages.BigDot2_png))
	handleError(err)
	bd.images[1], err = ebiten.NewImageFromImage(img2, ebiten.FilterDefault)
	handleError(err)
}

func (bd *bigDotManager) add(x,y int){
	bd.dots.PushBack(pos{x,y})
}

func (bd *bigDotManager) draw(sc *ebiten.Image){
	for e:= bd.dots.Front(); e != nil ; e = e.Next(){
		v := e.Value.(pos)
		x := float64(v.y * stageBlocSize)
		y := float64(v.x * stageBlocSize)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x,y)
		sc.DrawImage(bd.images[0], op)
	}
}