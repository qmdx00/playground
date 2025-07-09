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
