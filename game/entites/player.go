package entites

import (
	"math"

	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/M0rfes/a-go-game/game/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerSpeed           = 2.0
	playerNodeName        = "player"
	playerSpriteName      = "player_sprite"
	playerSpriteLayer     = 0
	playerWorldLayerIndex = 1
	playerRadius          = 6
)

type Player struct {
	*core.Node2D
	texture *core.Sprite
	tick    uint
	events  *PlayerEvents
}

type PlayerEvents struct {
	OnPlayerDead func()
}

func NewPlayer(input *core.InputManager, engine *core.Engine, event *PlayerEvents) *Player {
	world := engine.World()
	collisionManager := engine.CollisionManager()

	p := core.NewNode2D(playerNodeName)
	playerSprite := sprites.GetPlayerSprite(playerSpriteName)
	p.SetPosition(320, 240)
	p.AddChild(playerSprite)

	shape := core.NewCollisionCircle(playerRadius)
	mask := core.NewCollisionMask(core.LayerPlayer, core.LayerEnemy)
	c := collisionManager.NewCollider("player", shape, mask)
	c.SetPosition(0, 0)
	p.AddChild(c)

	c.SetOnCollide(func(other *core.Collider) {
		// end the game for now
		event.OnPlayerDead()
	})

	world.AddNodeToLayer(p, playerWorldLayerIndex)

	player := &Player{
		Node2D:  p,
		texture: playerSprite,
		tick:    0,
		events:  event,
	}

	input.AddAction("move_up", core.NewKeyRawInputButton(ebiten.KeyArrowUp))
	input.AddAction("move_up", core.NewKeyRawInputButton(ebiten.KeyW))
	input.AddAction("move_down", core.NewKeyRawInputButton(ebiten.KeyArrowDown))
	input.AddAction("move_down", core.NewKeyRawInputButton(ebiten.KeyS))
	input.AddAction("move_left", core.NewKeyRawInputButton(ebiten.KeyArrowLeft))
	input.AddAction("move_left", core.NewKeyRawInputButton(ebiten.KeyA))
	input.AddAction("move_right", core.NewKeyRawInputButton(ebiten.KeyArrowRight))
	input.AddAction("move_right", core.NewKeyRawInputButton(ebiten.KeyD))

	return player
}

func (p *Player) UpdateMovement(inp *core.InputManager) {
	p.tick++
	step := float64(p.tick) * 0.18
	bob := math.Abs(math.Sin(step)) * 6.0
	tilt := math.Sin(step) * 0.08

	p.texture.SetPosition(0, -bob)
	p.texture.SetRotation(tilt)
	dx, dy := 0.0, 0.0
	if inp.IsActionPressed("move_up") {
		dy -= 1
	}
	if inp.IsActionPressed("move_down") {
		dy += 1
	}
	if inp.IsActionPressed("move_left") {
		dx -= 1
	}
	if inp.IsActionPressed("move_right") {
		dx += 1
	}
	if dx == 0 && dy == 0 {
		return
	}
	length := math.Sqrt(dx*dx + dy*dy)
	if length > 0 {
		dx /= length
		dy /= length
	}
	x, y := p.GetPosition().X(), p.GetPosition().Y()
	p.SetPosition(x+dx*playerSpeed, y+dy*playerSpeed)
}
