package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const floorLayerIndex uint = 0

type Engine struct {
	world            *World
	input            *InputManager
	collisionManager *CollisionManager
}

func NewEngine() *Engine {
	return &Engine{
		world:            NewWorld(uint(Settings.ScreenWidth), uint(Settings.ScreenHeight)),
		input:            NewInputManager(),
		collisionManager: NewCollisionManager(),
	}
}

func (engine *Engine) World() *World {
	return engine.world
}

func (engine *Engine) CollisionManager() *CollisionManager {
	return engine.collisionManager
}

func (engine *Engine) Update() error {
	engine.input.Update()
	engine.world.Update()
	return nil
}

func (engine *Engine) Draw(target *ebiten.Image) {
	engine.world.Draw(target)
}

func (engine *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (engine *Engine) Input() *InputManager {
	return engine.input
}
