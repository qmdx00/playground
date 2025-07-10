package main

import "fmt"

type Transition struct {
	PlayerEvent PlayerEvent
	From, To    PlayerState
}

type FSM struct {
	currentState PlayerState
	transitions  map[PlayerEvent][]Transition
}

func NewPlayerFSM(initialState PlayerState) *FSM {
	return &FSM{
		currentState: initialState,
		transitions: map[PlayerEvent][]Transition{
			MoveRightEvent: {
				{From: PlayerStateIdle, To: PlayerStateRunning, PlayerEvent: MoveRightEvent},
				{From: PlayerStateRunning, To: PlayerStateRunning, PlayerEvent: MoveRightEvent},
			},
			StopEvent: {
				{From: PlayerStateRunning, To: PlayerStateIdle, PlayerEvent: StopEvent},
			},
		},
	}
}

func (f *FSM) Transition(PlayerEvent PlayerEvent) error {
	for _, t := range f.transitions[PlayerEvent] {
		if t.From == f.currentState {
			f.currentState = t.To
			return nil
		}
	}
	return fmt.Errorf("failed to transition from %s with event %s", f.currentState, PlayerEvent)
}
