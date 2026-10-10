package sprites

import (
	"github.com/M0rfes/a-go-game/assets"
	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

const FloorLayer uint = 0

// GetFloorTexture returns the default sand/dirt floor tile texture.
func GetFloorTexture() *ebiten.Image {
	return assets.Textures["sprites/tiny-dungeon/tile_0048.png"]
}

// GetStoneFloorTexture returns the stone floor tile texture.
func GetStoneFloorTexture() *ebiten.Image {
	return assets.Textures["sprites/tiny-dungeon/tile_0014.png"]
}

// NewFloorSprite constructs a configured floor tile sprite.
func NewFloorSprite(name string) *core.Sprite {
	tex := GetFloorTexture()
	sprite := core.NewSprite(name, tex, FloorLayer, false)
	sprite.SetPosition(0, 0)
	sprite.SetPivot(0, 0)
	b := tex.Bounds()
	tw, th := float64(b.Dx()), float64(b.Dy())
	if tw > 0 && th > 0 {
		sprite.SetScale(float64(core.Settings.ScreenWidth)/tw, float64(core.Settings.ScreenHeight)/th)
	}
	return sprite
}

// GetFloorSprite is a getter that constructs and returns a floor sprite.
func GetFloorSprite() *core.Sprite {
	return NewFloorSprite("floor")
}
