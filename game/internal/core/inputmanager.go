package core

type ActionBinding []RawInputButton

type InputManager struct {
	actions     map[string]ActionBinding
	actionMap   *ActionMap
	stateBuffer *StateBuffer
}

func NewInputManager() *InputManager {
	return &InputManager{
		actions:     make(map[string]ActionBinding),
		actionMap:   NewActionMap(),
		stateBuffer: NewStractBuffer(),
	}
}

func (i *InputManager) Update() {
	i.stateBuffer.Upadte()
}

func (i *InputManager) AddAction(action string, rawInputButton RawInputButton) {
	if _, exists := i.actions[action]; !exists {
		i.actions[action] = ActionBinding{}
	}
	i.actions[action] = append(i.actions[action], rawInputButton)
}

func (i *InputManager) Delete(action string) {
	delete(i.actions, action)
}

func (i *InputManager) IsActionPressed(actionName string) bool {
	buttons, exists := i.actions[actionName]
	if !exists {
		return false
	}

	for _, btn := range buttons {
		if btn.IsPressed() {
			return true
		}
	}

	return false
}

func (i *InputManager) IsActionJustPressed(actionName string) bool {
	buttons, exists := i.actions[actionName]
	if !exists {
		return false
	}

	for _, btn := range buttons {
		if btn.IsJustPressed() {
			return true
		}
	}

	return false
}

func (i *InputManager) IsActionJustReleased(actionName string) bool {
	buttons, exists := i.actions[actionName]
	if !exists {
		return false
	}

	for _, btn := range buttons {
		if btn.IsJustReleased() {
			return true
		}
	}

	return false
}

func (i *InputManager) IsActionReleased(actionName string) bool {
	buttons, exists := i.actions[actionName]
	if !exists {
		return true
	}

	for _, btn := range buttons {
		if btn.IsPressed() {
			return false
		}
	}

	return true
}

func (i *InputManager) RegisterAction(id ActionID, action Action) {
	i.actionMap.AddAction(id, action)
}

func (i *InputManager) IsActionActive(id ActionID) bool {
	action, ok := i.actionMap.GetAction(id)
	if !ok {
		return false
	}
	return i.stateBuffer.IsActionActive(action)
}

func (i *InputManager) ActionMap() *ActionMap {
	return i.actionMap
}
