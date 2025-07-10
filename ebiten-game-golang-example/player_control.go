package main

import "github.com/hajimehoshi/ebiten/v2"

var KeyboardEventMap = map[ebiten.Key]PlayerEvent{
	ebiten.KeyD:     MoveRightEvent,
	ebiten.KeyRight: MoveRightEvent,
}

func GetPlayerEventFromKey(key ebiten.Key) PlayerEvent {
	if event, ok := KeyboardEventMap[key]; ok {
		return event
	}
	return StopEvent
}
