package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func PrintTree(screen *ebiten.Image, posX, posY, distance, direction float64, depth int) {

	x := posX + distance*math.Sin(direction*math.Pi/180)
	y := posY - distance*math.Cos(direction*math.Pi/180)

	ebitenutil.DrawLine(screen, posX, posY, x, y, color.White)

	if depth > 0 {
		PrintTree(screen, x, y, distance*frac, direction-angle, depth-1)
		PrintTree(screen, x, y, distance*frac, direction+angle, depth-1)
	}
}
