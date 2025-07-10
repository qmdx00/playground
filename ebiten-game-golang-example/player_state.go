package main

import (
	"bytes"
	"ebiten-game-golang-example/images"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	frameOX     = 0
	frameOY     = 0
	frameWidth  = 32
	frameHeight = 32

	runFrameCount  = 8
	idleFrameCount = 9
	jumpFrameCount = 12
)

func init() {
	runImage, _, err := image.Decode(bytes.NewReader(images.Run_png))
	if err != nil {
		log.Fatal(err)
	}
	idleImage, _, err := image.Decode(bytes.NewReader(images.Idle_png))
	if err != nil {
		log.Fatal(err)
	}

	runSpriteImage = ebiten.NewImageFromImage(runImage)
	idleSpriteImage = ebiten.NewImageFromImage(idleImage)
}

type PlayerState string

const (
	PlayerStateIdle    PlayerState = "idle"
	PlayerStateRunning PlayerState = "running"
)

type PlayerEvent string

const (
	MoveRightEvent PlayerEvent = "move_right"
	StopEvent      PlayerEvent = "stop"
)

func (state PlayerState) Image(frame int) *ebiten.Image {
	switch state {
	case PlayerStateRunning:
		index := (frame / 10) % runFrameCount
		sx, sy := frameOX+index*frameWidth, frameOY
		return runSpriteImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
	case PlayerStateIdle:
		index := (frame / 10) % idleFrameCount
		sx, sy := frameOX+index*frameWidth, frameOY
		return idleSpriteImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
	}
	return nil
}
