package main

import "github.com/hajimehoshi/ebiten/v2"

var KeyboardEventMap = map[ebiten.Key]PlayerEvent{
	ebiten.KeyA:    MoveLeftEvent,
	ebiten.KeyLeft: MoveLeftEvent,

	ebiten.KeyD:     MoveRightEvent,
	ebiten.KeyRight: MoveRightEvent,

	ebiten.KeyS:    MoveDownEvent,
	ebiten.KeyDown: MoveDownEvent,

	ebiten.KeyW:  MoveUpEvent,
	ebiten.KeyUp: MoveUpEvent,
}

func GetPlayerEventFromKey(key ebiten.Key) PlayerEvent {
	if event, ok := KeyboardEventMap[key]; ok {
		return event
	}
	return StopEvent
}
