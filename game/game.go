package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"bytes"
	"image"
	_ "image/png"

	"github.com/M0rfes/a-go-game/assets"
	"github.com/M0rfes/a-go-game/game/internal/core"
)

const rotationSpeed = 0.02

type Game struct {
	engine *core.Engine
	logo   *core.Node2D
}

func (game *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return game.engine.Layout(outsideWidth, outsideHeight)
}

func (game *Game) Update() error {
	game.logo.SetRotation(game.logo.GetRotation() + rotationSpeed)
	if game.logo.GetRotation() >= 2*math.Pi {
		game.logo.SetRotation(0)
	}
	return game.engine.Update()
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.engine.Draw(screen)
}

func Run() error {
	img, _, err := image.Decode(bytes.NewReader(assets.Images["hero/walk/down/00.png"]))
	if err != nil {
		return err
	}
	texture := ebiten.NewImageFromImage(img)
	engine := core.NewEngine()
	world := engine.World()

	logo := core.NewNode2D("logo")
	logo.SetPosition(320, 240)
	sprite := core.NewSprite("hero", texture, 0, true)
	sprite.SetPosition(0, 0)
	logo.AddChild(sprite)
	world.AddNodeToDefaultLayer(logo)
	ebiten.SetWindowSize(Settings.ScreenWidth, Settings.ScreenHeight)
	ebiten.SetWindowTitle("A go game")
	game := &Game{engine: engine, logo: logo}
	return ebiten.RunGame(game)
}
