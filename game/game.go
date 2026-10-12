package game

import (
	"github.com/M0rfes/a-go-game/game/entites"
	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/M0rfes/a-go-game/game/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine *core.Engine
	player *entites.Player
}

func (game *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return game.engine.Layout(outsideWidth, outsideHeight)
}

func (game *Game) Update() error {

	// Classic retro walk-in-place animation (bob up/down + slight side tilt):

	input := game.engine.Input()
	game.player.UpdateMovement(input)
	return game.engine.Update()
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.engine.Draw(screen)
}

func Run() error {
	engine := core.NewEngine()
	world := engine.World()
	input := engine.Input()
	floor := sprites.GetFloorSprite()

	world.AddNodeToLayer(floor, sprites.FloorLayer)

	player := entites.NewPlayer(input, world)
	ebiten.SetWindowSize(core.Settings.ScreenWidth, core.Settings.ScreenHeight)
	ebiten.SetWindowTitle("A go game")

	game := &Game{
		engine: engine,
		player: player,
	}
	return ebiten.RunGame(game)
}
