package pacman

import (
	"bytes"
	"container/list"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)
type dotManager struct {
	dots *list.List
	image *ebiten.Image
}

func newDotManager() *dotManager {
	d := &dotManager{}
	d.dots = list.New()
	d.loadImage()
	return d
}

func (d *dotManager) loadImage() {
	img, _, err := image.Decode(bytes.NewReader(pacimages.Dot_png))
	handleError(err)
	d.image, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handleError(err)
}

func (d *dotManager) add(x,y int){
	d.dots.PushBack(pos{x,y})
}

func (d *dotManager) draw(sc *ebiten.Image){
	for e:= d.dots.Front(); e != nil ; e = e.Next(){
		v := e.Value.(pos)
		x := float64(v.y * stageBlocSize)
		y := float64(v.x * stageBlocSize)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x,y)
		sc.DrawImage(d.image, op)
	}
}