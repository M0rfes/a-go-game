package game

import (
	"math"

	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/M0rfes/a-go-game/game/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine     *core.Engine
	heroSprite *core.Sprite
	tick       int
}

func (game *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return game.engine.Layout(outsideWidth, outsideHeight)
}

func (game *Game) Update() error {
	game.tick++

	// Classic retro walk-in-place animation (bob up/down + slight side tilt):
	step := float64(game.tick) * 0.18
	bob := math.Abs(math.Sin(step)) * 6.0
	tilt := math.Sin(step) * 0.08

	game.heroSprite.SetPosition(0, -bob)
	game.heroSprite.SetRotation(tilt)

	return game.engine.Update()
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.engine.Draw(screen)
}

func Run() error {
	engine := core.NewEngine()
	world := engine.World()
	floor := sprites.GetFloorSprite()

	p := core.NewNode2D("player_shell")
	player := sprites.GetPlayerSprite()
	p.SetPosition(320, 240)
	p.AddChild(player)

	world.AddNodeToLayer(floor, sprites.FloorLayer)
	world.AddNodeToLayer(p, sprites.PlayerLayer)

	ebiten.SetWindowSize(core.Settings.ScreenWidth, core.Settings.ScreenHeight)
	ebiten.SetWindowTitle("A go game")

	game := &Game{
		engine:     engine,
		heroSprite: player,
	}
	return ebiten.RunGame(game)
}
