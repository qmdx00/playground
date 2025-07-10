package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneMap struct {
	tileLayers [screenWidth][screenHeight]string
}

func getDefaultTileLayer() [screenWidth][screenHeight]string {
	var layer [screenWidth][screenHeight]string

	// layer[15][10] = "1.png"
	// layer[15][11] = "2.png"
	// layer[15][12] = "2.png"
	// layer[15][13] = "2.png"
	// layer[15][14] = "3.png"

	// layer[16][10] = "4.png"
	// layer[16][11] = "5.png"
	// layer[16][12] = "5.png"
	// layer[16][13] = "5.png"
	// layer[16][14] = "6.png"

	for w := range screenWidth / unitWidth {
		layer[22][w] = "14.png"
	}

	return layer
}

func NewDefaultSceneMap() *SceneMap {
	return &SceneMap{tileLayers: getDefaultTileLayer()}
}

func (s *SceneMap) GetTileAt(x, y int) *Tile {
	if x < 0 || y < 0 || y >= len(s.tileLayers) || x >= len(s.tileLayers[y]) {
		return nil
	}
	return tileSet[s.tileLayers[y][x]]
}

func (s *SceneMap) Render(screen *ebiten.Image) {
	for y, row := range s.tileLayers {
		for x, tileID := range row {
			if tileID != "" {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(unitWidth)/tileWidth*tileScale, float64(unitHeight)/tileHeight*tileScale)
				op.GeoM.Translate(float64(x*unitWidth), float64(y*unitHeight))
				if tileImage, ok := tileSet[tileID]; ok && tileImage != nil {
					screen.DrawImage(tileImage.Image, op)
				}
			}
		}
	}
}
