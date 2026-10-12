package core

import "github.com/hajimehoshi/ebiten/v2"

type ActionID int

type ActionMap struct {
	actions map[ActionID]Action
}

func NewActionMap() *ActionMap {
	return &ActionMap{
		actions: make(map[ActionID]Action),
	}
}

func (am *ActionMap) AddAction(actionId ActionID, action Action) {
	am.actions[actionId] = action
}

func (am *ActionMap) RemoveAction(id ActionID) {
	delete(am.actions, id)
}

func (am *ActionMap) GetAction(id ActionID) (Action, bool) {
	action, ok := am.actions[id]
	return action, ok
}

func (am *ActionMap) ClearAction() {
	am.actions = make(map[ActionID]Action)
}

type ActionNotFoundError struct{}

func (e *ActionNotFoundError) Error() string {
	return "action not found"
}

var ErrActionNotFound = &ActionNotFoundError{}

func (am *ActionMap) SetKeyBinding(id ActionID, key ebiten.Key) error {
	action, exists := am.actions[id]
	if !exists {
		return ErrActionNotFound
	}
	am.actions[id] = NewKeyAction(key, action.GetMode())
	return nil
}
