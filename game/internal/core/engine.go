package core

import "github.com/hajimehoshi/ebiten/v2"

type Engine struct {
	world *World
}

func NewEngine() *Engine {
	return &Engine{world: NewWorld()}
}

func (engine *Engine) World() *World {
	return engine.world
}

func (engine *Engine) Update() error {
	engine.world.Update()
	return nil
}

func (engine *Engine) Draw(target *ebiten.Image) {
	engine.world.Draw(target)
}

func (engine *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}
