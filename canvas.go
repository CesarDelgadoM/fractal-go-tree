package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	length = 30
	frac   = 1
)

var angle float64 = 1
var depth int = 1

type Canvas struct {
	width  int
	height int
}

func NewCanvas(screenWidth int, screenHeight int) *Canvas {
	return &Canvas{
		width:  screenWidth,
		height: screenHeight,
	}
}

func (c *Canvas) Layout(width int, height int) (int, int) {
	return width, height
}

func (c *Canvas) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && depth <= 15 {
		depth++
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && depth >= 1 {
		depth--
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && angle <= 20 {
		angle++
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && angle >= 1 {
		angle--
	}

	return nil
}

func (c *Canvas) Draw(screen *ebiten.Image) {

	msg := fmt.Sprint("Presione las teclas flecha izquierda o derecha para graduar el angulo\n" +
		"Presione las teclas flecha arriba o abajo para graduar la profundidad")
	ebitenutil.DebugPrint(screen, msg)

	PrintTree(screen, screenWidth/2, screenHeight*9/10, length, 0, depth)

}
