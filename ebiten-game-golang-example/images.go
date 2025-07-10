package main

import (
	"bytes"
	"ebiten-game-golang-example/images"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	runSpriteImage  *ebiten.Image
	idleSpriteImage *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	runImage, _, err := image.Decode(bytes.NewReader(images.Run_png))
	if err != nil {
		log.Fatal(err)
	}
	runSpriteImage = ebiten.NewImageFromImage(runImage)

	idleImage, _, err := image.Decode(bytes.NewReader(images.Idle_png))
	if err != nil {
		log.Fatal(err)
	}
	idleSpriteImage = ebiten.NewImageFromImage(idleImage)

	backgroundImg, _, err := image.Decode(bytes.NewReader(images.Background_png))
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage = ebiten.NewImageFromImage(backgroundImg)
}
