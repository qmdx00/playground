package images

import (
	_ "embed"
)

// Sprite images download from:
// https://www.piskelapp.com/

var (
	//go:embed run.png
	Run_png []byte

	//go:embed idle.png
	Idle_png []byte
)

// Tileset images download from:
// https://www.gameart2d.com/free-platformer-game-tileset.html

var (
	//go:embed freetileset/Background/BG.png
	Background_png []byte
)
