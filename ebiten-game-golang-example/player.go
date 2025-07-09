package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y          float64
	width, height int
	step          float64

	frame int
	fsm   *FSM
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X: x, Y: y,
		width: 32, height: 32,
		step:  1.0,
		frame: 0,
		fsm:   NewPlayerFSM(PlayerStateIdle),
	}
}

func (p *Player) Update() {
	// make player animation
	// This function can be used to update player state,
	// e.g., animation frames.
	p.frame++
}

func (p *Player) Transition(g *Game, event PlayerEvent) {
	var dx, dy float64
	switch event {
	case MoveLeftEvent:
		dx = -p.step
	case MoveRightEvent:
		dx = p.step
	case MoveUpEvent:
		dy = -p.step
	case MoveDownEvent:
		dy = p.step
	case StopEvent:
	default:
	}

	// change player state
	p.fsm.Transition(event)

	// update player position
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

func (p *Player) RenderImage() *ebiten.Image {
	return p.fsm.currentState.Image(p.frame)
}
