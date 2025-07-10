package main

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	ID    string
	Image *ebiten.Image
}

func NewTile(id string, image *ebiten.Image) *Tile {
	return &Tile{
		ID:    id,
		Image: image,
	}
}
