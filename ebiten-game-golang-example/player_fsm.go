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
			MoveLeftEvent: {
				{From: PlayerStateIdle, To: PlayerStateWalking, PlayerEvent: MoveLeftEvent},
				{From: PlayerStateWalking, To: PlayerStateWalking, PlayerEvent: MoveLeftEvent},
			},
			MoveRightEvent: {
				{From: PlayerStateIdle, To: PlayerStateWalking, PlayerEvent: MoveRightEvent},
				{From: PlayerStateWalking, To: PlayerStateWalking, PlayerEvent: MoveRightEvent},
			},
			MoveDownEvent: {
				{From: PlayerStateIdle, To: PlayerStateWalking, PlayerEvent: MoveDownEvent},
				{From: PlayerStateWalking, To: PlayerStateWalking, PlayerEvent: MoveDownEvent},
			},
			MoveUpEvent: {
				{From: PlayerStateIdle, To: PlayerStateWalking, PlayerEvent: MoveUpEvent},
				{From: PlayerStateWalking, To: PlayerStateWalking, PlayerEvent: MoveUpEvent},
			},
			StopEvent: {
				{From: PlayerStateWalking, To: PlayerStateIdle, PlayerEvent: StopEvent},
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
