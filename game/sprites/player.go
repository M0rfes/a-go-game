package sprites

import (
	"github.com/M0rfes/a-go-game/assets"
	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

const PlayerLayer uint = 1

// GetPlayerTexture returns the player texture from assets.
func GetPlayerTexture() *ebiten.Image {
	return assets.Textures["sprites/tiny-dungeon/tile_0085.png"]
}

// NewPlayerSprite constructs a configured player sprite.
func NewPlayerSprite(name string) *core.Sprite {
	tex := GetPlayerTexture()
	sprite := core.NewSprite(name, tex, PlayerLayer, true)
	return sprite
}

// GetPlayerSprite is a getter that constructs and returns a player sprite.
func GetPlayerSprite(name string) *core.Sprite {
	return NewPlayerSprite(name)
}
