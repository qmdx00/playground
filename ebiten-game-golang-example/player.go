package main

import (
	"bytes"
	"ebiten-game-golang-example/images"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Player struct {
	X, Y          float64
	width, height int
	image         *ebiten.Image
	frame         int
	step          float64
}

func NewPlayer(x, y float64) *Player {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		X:      x,
		Y:      y,
		width:  32,
		height: 32,
		image:  ebiten.NewImageFromImage(img),
		step:   1.0, // step for animation
	}
}

func (p *Player) Update() {
	// make player animation
	// This function can be used to update player state,
	// e.g., animation frames.

	p.frame++
}

func (p *Player) Move(g *Game, direction Direction) {
	var dx, dy float64

	switch direction {
	case Left:
		dx = -p.step
	case Right:
		dx = p.step
	case Up:
		dy = -p.step
	case Down:
		dy = p.step
	}

	p.X += dx
	p.Y += dy

	// Ensure the player stays within the game boundaries
	if p.X < 0 {
		p.X = 0
	} else if p.X+float64(p.width) > float64(g.width) {
		p.X = float64(g.width) - float64(p.width)
	}

	if p.Y < 0 {
		p.Y = 0
	} else if p.Y+float64(p.height) > float64(g.height) {
		p.Y = float64(g.height) - float64(p.height)
	}
}

func (p *Player) Image() *ebiten.Image {
	index := (p.frame / 5) % runnerFrameCount
	sx, sy := 0, index*p.height
	// fmt.Println("Player frame index:", index, "sx:", sx, "sy:", sy)
	return p.image.SubImage(image.Rect(sx, sy, sx+p.width, sy+p.height)).(*ebiten.Image)
}
