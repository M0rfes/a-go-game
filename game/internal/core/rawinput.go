package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputButtonType uint

const (
	ButtonKeyboard InputButtonType = iota
	ButtonMouse
	ButtonGamepad
)

type RawInputButton struct {
	key         ebiten.Key
	mouseButton ebiten.MouseButton
	gamepadID   ebiten.GamepadID
	gamepadBtn  ebiten.GamepadButton
	buttonType  InputButtonType
}

func NewKeyInputButton(key ebiten.Key) *RawInputButton {
	return &RawInputButton{
		key:        key,
		buttonType: ButtonKeyboard,
	}
}

func NewMouseRawInputButton(btn ebiten.MouseButton) *RawInputButton {
	return &RawInputButton{
		mouseButton: btn,
		buttonType:  ButtonMouse,
	}
}

func NewGamepadRawInputButton(id ebiten.GamepadID, btn ebiten.GamepadButton) *RawInputButton {
	return &RawInputButton{
		gamepadID:  id,
		gamepadBtn: btn,
		buttonType: ButtonGamepad,
	}
}

func (r *RawInputButton) IsPressed() bool {
	switch r.buttonType {
	case ButtonKeyboard:
		return ebiten.IsKeyPressed(r.key)
	case ButtonMouse:
		return ebiten.IsMouseButtonPressed(r.mouseButton)
	case ButtonGamepad:
		return ebiten.IsGamepadButtonPressed(r.gamepadID, r.gamepadBtn)
	}
	return false
}

func (r *RawInputButton) IsJustPressed() bool {
	switch r.buttonType {
	case ButtonKeyboard:
		return inpututil.IsKeyJustPressed(r.key)
	case ButtonMouse:
		return inpututil.IsMouseButtonJustPressed(r.mouseButton)
	case ButtonGamepad:
		return inpututil.IsGamepadButtonJustPressed(r.gamepadID, r.gamepadBtn)
	}
	return false
}

func (r *RawInputButton) IsJustReleased() bool {
	switch r.buttonType {
	case ButtonKeyboard:
		return inpututil.IsKeyJustReleased(r.key)
	case ButtonMouse:
		return inpututil.IsMouseButtonJustReleased(r.mouseButton)
	case ButtonGamepad:
		return inpututil.IsGamepadButtonJustReleased(r.gamepadID, r.gamepadBtn)
	}
	return false
}

func (r *RawInputButton) PressDuration() int {
	switch r.buttonType {
	case ButtonKeyboard:
		return inpututil.KeyPressDuration(r.key)
	case ButtonMouse:
		return inpututil.MouseButtonPressDuration(r.mouseButton)
	case ButtonGamepad:
		return inpututil.GamepadButtonPressDuration(r.gamepadID, r.gamepadBtn)
	}
	return 0
}

func NewKeyRawInputButton(key ebiten.Key) RawInputButton {
	return RawInputButton{
		key:        key,
		buttonType: ButtonKeyboard,
	}
}
