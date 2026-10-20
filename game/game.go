package game

import (
	"github.com/M0rfes/a-go-game/game/entites"
	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine    *core.Engine
	player    *entites.Player
	enemy     *entites.Enemy
	hud       *HUD
	gameOver  bool
	gameStart bool
}

func (game *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return game.engine.Layout(outsideWidth, outsideHeight)
}

func (game *Game) Update() error {
	input := game.engine.Input()
	game.hud.update(input)
	if !game.gameStart {
		return game.engine.Update()
	}
	if game.gameOver {
		return game.engine.Update()
	}

	// Classic retro walk-in-place animation (bob up/down + slight side tilt):

	game.player.UpdateMovement(input)
	game.enemy.FollowPlayer(game.player)

	game.engine.CollisionManager().CheckCollision()
	return game.engine.Update()
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.engine.Draw(screen)
	game.hud.draw(screen)

	if game.gameOver {
		game.hud.gameOver(screen)
	}
}

func Run() error {
	engine := core.NewEngine()
	world := engine.World()
	input := engine.Input()
	game := &Game{
		engine: engine,
	}

	entites.NewTileMap(world)

	hud := NewHUD(input, &HUDEvens{
		StartGame: func() {
			game.gameStart = true
		},
	})

	game.hud = hud

	player := entites.NewPlayer(input, engine, &entites.PlayerEvents{
		OnPlayerDead: func() {
			game.gameOver = true
		},
	})
	game.player = player
	game.enemy = entites.NewEnemy(engine, 320, 240)
	ebiten.SetWindowSize(core.Settings.ScreenWidth, core.Settings.ScreenHeight)
	ebiten.SetWindowTitle("A go game")
	world.Camera().SetFollow(player)

	return ebiten.RunGame(game)
}
