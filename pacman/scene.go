package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
)


type scene struct{
	matrix 		[][]elem
	wallSurface *ebiten.Image
	images      map[elem]*ebiten.Image
	stage		*stage
	dotManager  *dotManager
	bigDotManager  *bigDotManager
}

func newScene(st *stage) *scene{
	s := &scene{}
	s.stage = st;
	if s.stage == nil{
		s.stage = defaultStage
	}
	s.images = make(map[elem]*ebiten.Image)
	s.dotManager = newDotManager()
	s.bigDotManager = newBigDotManager()
	s.loadImages()
	s.createStage()
	s.buildWallSurface()
	return s
}

func (scene *scene) ScreenWidth() int {
	return len(scene.stage.matrix[0])*stageBlocSize
}

func (scene *scene) ScreenHeight() int {
	return len(scene.stage.matrix)*stageBlocSize
}

func (s *scene) createStage() {
	w := len( s.stage.matrix[0] )
	h := len( s.stage.matrix )
	s.matrix = make([][]elem, h)
	for i:=0 ; i < h; i++ {
		s.matrix[i] = make([]elem, w)
		for j:=0; j < w; j++ {
			c := s.stage.matrix[i][j] - '0'
			if c <= 9 {
				s.matrix[i][j] = elem(c)
			} else{
				s.matrix[i][j] = elem(s.stage.matrix[i][j] - 'a' + 10 )
			}

			switch s.matrix[i][j] {
			case dotElem:
				s.dotManager.add(i,j)
			case bigDotElem:
				s.bigDotManager.add(i,j)
			}
		}
	}
}


func (s *scene) buildWallSurface(){
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])

	sizeH := ((h*stageBlocSize)/backgroundImageSize + 1 ) * backgroundImageSize
	sizeW := ((w*stageBlocSize)/backgroundImageSize + 1 ) * backgroundImageSize

	s.wallSurface, _ = ebiten.NewImage(sizeW, sizeH, ebiten.FilterDefault)

	for i:=0; i < sizeH/backgroundImageSize; i++ {
		y := float64(i * backgroundImageSize)
		for j:=0; j < sizeW/backgroundImageSize; j++ {
			x := float64(j * backgroundImageSize)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x,y)
			s.wallSurface.DrawImage(s.images[backgroundElem], op)
		}
	}

	for i:=0; i < h; i++ {
		y := float64( i * stageBlocSize)
		for j:=0; j < w; j++ {
			if !isWall(s.matrix[i][j]) {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			x := float64( j * stageBlocSize)
			op.GeoM.Translate(x,y)
			s.wallSurface.DrawImage(s.images[s.matrix[i][j]], op)
		}
	}
}

func (s *scene) loadImages() {
	for i := w0; i <= w24; i++ {
		img, _, err := image.Decode(bytes.NewReader(pacimages.WallImages[i]))
		handleError(err)
		s.images[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	img, _, err := image.Decode(bytes.NewReader(pacimages.Background_png))
	handleError(err)
	s.images[backgroundElem], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handleError(err)
}

func (s *scene) update(screen *ebiten.Image) error {
	if(ebiten.IsDrawingSkipped()){
		return nil
	}
	screen.Clear()
	screen.DrawImage(s.wallSurface, nil)
	s.dotManager.draw(screen)
	s.bigDotManager.draw(screen)
	return nil
}
