package core

import "github.com/hajimehoshi/ebiten/v2"

type InputState struct {
	IsPressed  bool
	wasPressed bool
}

type StateBuffer struct {
	keyStates   map[ebiten.Key]InputState
	mouseStates map[ebiten.MouseButton]InputState
}

func NewStractBuffer() *StateBuffer {
	return &StateBuffer{
		keyStates:   make(map[ebiten.Key]InputState),
		mouseStates: make(map[ebiten.MouseButton]InputState),
	}
}

func (s *StateBuffer) Upadte() {
	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		old := s.keyStates[key]
		s.keyStates[key] = InputState{
			IsPressed:  ebiten.IsKeyPressed(key),
			wasPressed: old.IsPressed,
		}
	}

	for btn := ebiten.MouseButtonLeft; btn <= ebiten.MouseButtonRight; btn++ {
		old := s.mouseStates[btn]
		s.mouseStates[btn] = InputState{
			IsPressed:  ebiten.IsMouseButtonPressed(btn),
			wasPressed: old.IsPressed,
		}
	}
}

func (s *StateBuffer) checkState(state InputState, mode ActionMode) bool {
	if mode.Has(ActionHold) && state.IsPressed {
		return true
	}

	if mode.Has(ActionPressOnce) && state.IsPressed && !state.wasPressed {
		return true
	}

	if mode.Has(ActionReleaseOnce) && !state.IsPressed && state.wasPressed {
		return true
	}
	return false
}

func (s *StateBuffer) IsActionActive(action Action) bool {
	switch action.GetActionType() {
	case KeyAction:
		state := s.keyStates[action.GetKey()]
		return s.checkState(state, action.GetMode())
	case MouseButtonAction:
		state := s.mouseStates[action.GetMouseButton()]
		return s.checkState(state, action.GetMode())
	}
	return false
}
