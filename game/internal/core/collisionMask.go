package core

type collisionID uint

const (
	LayerPlayer collisionID = 1 << iota
	LayerEnemy
)

type CollisionMask struct {
	identity     collisionID
	collidesWith collisionID
}

func NewCollisionMask(identity, collidesWith collisionID) CollisionMask {
	return CollisionMask{
		identity:     identity,
		collidesWith: collidesWith,
	}
}

func (m CollisionMask) CanCollideWith(other CollisionMask) bool {
	return (m.collidesWith & other.identity) != 0
}

func (m CollisionMask) GetIdentity() collisionID {
	return m.identity
}
