package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/guru056/pacmango/pacman"
)

func main() {
	g := pacman.NewGame()
	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
