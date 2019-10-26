package pacman

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/kgosse/pacmanresources/images"
	"image"
)

//struct ghostManager -> ghosts , images, vulnerabilityImages
//constructor ghostManager
//loadImages
//loadPacmanImages
//loadVulnerabilityImages
//addGhost
//draw

type ghostManager struct {
	ghosts				[]*ghost
	images				map[elem][8]*ebiten.Image
	vulnerabilityImages [5]*ebiten.Image
}

func newGhostManager() *ghostManager{
	gm := &ghostManager{};
	gm.images	=	make(map[elem][8]*ebiten.Image)
	gm.loadImages()
	return gm
}

func (gm *ghostManager) addGhost(x,y int , kind elem) {
	gm.ghosts = append(gm.ghosts, newGhost(x,y,kind))
}

func (gm *ghostManager) loadImages() {
	gm.images[blinkyElem] =	gm.loadGhostImages(pacimages.BlinkyImages)
	gm.images[clydeElem] =	gm.loadGhostImages(pacimages.ClydeImages)
	gm.images[inkyElem] =	gm.loadGhostImages(pacimages.InkyImages)
	gm.images[pinkyElem] =	gm.loadGhostImages(pacimages.PinkyImages)
	gm.vulnerabilityImages =	gm.loadVulnerabilityImages(pacimages.VulnerabilityImages)

}

func (gm *ghostManager) loadGhostImages(g [8][]byte) [8]*ebiten.Image{
	var ghostImgs [8]*ebiten.Image
	for i:=0; i<8; i++ {
		img, _, err := image.Decode(bytes.NewReader(g[i]))
		handleError(err)
		ghostImgs[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	return ghostImgs
}

func (gm *ghostManager) loadVulnerabilityImages(g [5][]byte) [5]*ebiten.Image{
	var vulnerabilityImgs [5]*ebiten.Image
	for i:=0; i<5; i++ {
		img, _, err := image.Decode(bytes.NewReader(g[i]))
		handleError(err)
		vulnerabilityImgs[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	return vulnerabilityImgs
}

func (gm *ghostManager) draw(screen *ebiten.Image) {
	for i:=0; i<len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		imgs,_ 	:= gm.images[g.kind]
		images  := make([]*ebiten.Image , 13)
		copy(images, imgs[:])
		copy(images[8:], gm.vulnerabilityImages[:])
		g.draw(screen, images)
	}
}