package sprites

import (
	"github.com/M0rfes/a-go-game/assets"
	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

const EnemyLayer uint = 2

func GetEnemyTexture() *ebiten.Image {
	return assets.Textures["sprites/tiny-dungeon/tile_0096.png"]
}

func GetEnemySprite(name string) *core.Sprite {
	tex := GetEnemyTexture()
	return core.NewSprite(name, tex, EnemyLayer, true)
}
