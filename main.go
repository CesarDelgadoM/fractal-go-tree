package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	canvas := NewCanvas(screenWidth, screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Fractal-go-tree")

	if err := ebiten.RunGame(canvas); err != nil {
		panic(err)
	}
}
