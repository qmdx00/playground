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

	tileSet = make(map[string]*Tile)
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

	if err := loadTiles(); err != nil {
		log.Fatal(err)
	}
}

func loadTiles() error {
	path := "freetileset/Tiles"
	tiles, err := images.Tiles_png.ReadDir(path)
	if err != nil {
		return err
	}

	for _, tile := range tiles {
		if tile.IsDir() {
			continue
		}

		imgData, err := images.Tiles_png.ReadFile(path + "/" + tile.Name())
		if err != nil {
			return err
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			return err
		}

		tileImage := ebiten.NewImageFromImage(img)
		tileSet[tile.Name()] = NewTile(tile.Name(), tileImage)
	}
	return nil
}
