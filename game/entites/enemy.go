package entites

import (
	"math"

	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/M0rfes/a-go-game/game/sprites"
)

const (
	enemyNodeName        = "enemy"
	enemySpriteName      = "enemy_sprite"
	enemyWorldLayerIndex = 1
	enemyRadius          = 6
	enemySpeed           = 0.5
	enemySpawnOffset     = 120.0
)

type Enemy struct {
	*core.Node2D
}

func NewEnemy(engine *core.Engine, playerStartX, playerStartY float64) *Enemy {
	world := engine.World()
	collisionManager := engine.CollisionManager()

	e := core.NewNode2D(enemyNodeName)
	e.AddChild(sprites.GetEnemySprite(enemySpriteName))
	e.SetPosition(playerStartX+enemySpawnOffset, playerStartY)

	shape := core.NewCollisionCircle(enemyRadius)
	mask := core.NewCollisionMask(core.LayerEnemy, core.LayerPlayer)
	c := collisionManager.NewCollider("enemy", shape, mask)
	c.SetPosition(0, 0)
	e.AddChild(c)

	world.AddNodeToLayer(e, enemyWorldLayerIndex)

	return &Enemy{Node2D: e}
}

func (e *Enemy) Update() {
	// Follow target is set from game loop via FollowPlayer.
}

func (e *Enemy) FollowPlayer(player *Player) {
	if player == nil {
		return
	}
	px, py := player.GetPosition().X(), player.GetPosition().Y()
	ex, ey := e.GetPosition().X(), e.GetPosition().Y()
	dx, dy := px-ex, py-ey
	length := math.Sqrt(dx*dx + dy*dy)
	if length < 0.001 {
		return
	}
	dx /= length
	dy /= length
	e.SetPosition(ex+dx*enemySpeed, ey+dy*enemySpeed)
}
