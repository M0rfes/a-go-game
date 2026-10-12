package core

import "github.com/hajimehoshi/ebiten/v2"

type ActionMode uint

const (
	ActionHold ActionMode = 1 << iota
	ActionPressOnce
	ActionReleaseOnce
)

func (a ActionMode) Has(flag ActionMode) bool {
	return a&flag != 0
}

type ActionType int

const (
	KeyAction ActionType = iota
	MouseButtonAction
	GamepadButtonAction
)

type Action struct {
	actionType    ActionType
	mode          ActionMode
	key           ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadID     ebiten.GamepadID
	gamepadButton ebiten.GamepadButton
}

func NewKeyAction(key ebiten.Key, mode ActionMode) Action {
	return Action{
		actionType: KeyAction,
		mode:       mode,
		key:        key,
	}
}

func NewMouseAction(mouseButton ebiten.MouseButton, mode ActionMode) Action {
	return Action{
		actionType:  MouseButtonAction,
		mode:        mode,
		mouseButton: mouseButton,
	}
}

func NewGamepadButton(gamepadID ebiten.GamepadID, gamepadeButton ebiten.GamepadButton, mode ActionMode) Action {
	return Action{
		actionType:    GamepadButtonAction,
		gamepadID:     gamepadID,
		gamepadButton: gamepadeButton,
		mode:          mode,
	}
}

func (a *Action) GetActionType() ActionType {
	return a.actionType
}

func (a *Action) GetMode() ActionMode {
	return a.mode
}

func (a *Action) GetKey() ebiten.Key {
	return a.key
}

func (a *Action) GetMouseButton() ebiten.MouseButton {
	return a.mouseButton
}

func (a *Action) GetGamepadButton() (ebiten.GamepadID, ebiten.GamepadButton) {
	return a.gamepadID, a.gamepadButton
}
