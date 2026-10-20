package game

import (
	"image/color"

	"github.com/M0rfes/a-go-game/game/internal/core"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type HUDEvens struct {
	StartGame func()
}

type HUD struct {
	gameOverLabel *ebiten.Image
	events        *HUDEvens
}

func NewHUD(input *core.InputManager, events *HUDEvens) *HUD {
	face := text.NewGoXFace(basicfont.Face7x13)
	w, h := text.Measure("GAME OVER", face, 0)
	img := ebiten.NewImage(int(w)+2, int(h)+2)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(1, 1)
	opts.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	text.Draw(img, "GAME OVER", face, opts)

	input.AddAction("start_game", core.NewKeyRawInputButton(ebiten.KeySpace))
	return &HUD{
		gameOverLabel: img,
		events:        events,
	}
}

func (h *HUD) draw(screen *ebiten.Image) {
}

func (h *HUD) update(input *core.InputManager) {
	if input.IsActionPressed("start_game") {
		h.events.StartGame()
	}
}

func (hud *HUD) gameOver(screen *ebiten.Image) {
	if hud.gameOverLabel == nil {
		return
	}

	b := hud.gameOverLabel.Bounds()

	iw := float64(b.Dx())
	ih := float64(b.Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-iw/2, -ih/2)
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(float64(core.Settings.ScreenWidth)/2, float64(core.Settings.ScreenHeight)/2)
	screen.DrawImage(hud.gameOverLabel, op)
}
