package core

type CollisionManager struct {
	colliders []*Collider
}

func NewCollisionManager() *CollisionManager {
	return &CollisionManager{
		colliders: make([]*Collider, 0),
	}
}

func (m *CollisionManager) AddCollider(c *Collider) {
	if c == nil {
		return
	}

	m.colliders = append(m.colliders, c)
}

func (m *CollisionManager) NewCollider(name string, shape *CollisionCircul, mask CollisionMask) *Collider {
	if m == nil || shape == nil {
		return nil
	}

	c := &Collider{
		Node2D: NewNode2D(name),
		shape:  shape,
		mask:   mask,
	}

	m.AddCollider(c)
	return c
}

type circleSample struct {
	x, y, r float64
}

func sampleCircle(c *Collider) circleSample {
	pos := c.GetWorldPosition()
	return circleSample{
		x: pos.X(),
		y: pos.Y(),
		r: c.shape.GetRadius(),
	}
}

func (m *CollisionManager) CheckCollision() {
	for i := 0; i < len(m.colliders); i++ {
		a := m.colliders[i]
		if a == nil {
			continue
		}

		sampleA := sampleCircle(a)
		for j := i + 1; j < len(m.colliders); j++ {
			b := m.colliders[j]
			if b == nil {
				continue
			}
			if !pairWantsCollision(a, b) {
				continue
			}
			checkColliderPair(a, sampleA, b)
		}
	}
}

func checkColliderPair(a *Collider, sampleA circleSample, b *Collider) {
	if a == nil || b == nil {
		return
	}

	if !pairWantsCollision(a, b) {
		return
	}

	if !samplesOverlap(sampleA, sampleCircle(b)) {
		return
	}

	notifyPair(a, b)
}

func pairWantsCollision(a, b *Collider) bool {
	return a.CanCollideWith(b) && b.CanCollideWith(a)
}

func samplesOverlap(a, b circleSample) bool {
	return circlesOverlap(a.x, a.y, a.r, b.x, b.y, b.r)
}

func notifyPair(a, b *Collider) {
	if a.onCollide != nil {
		a.onCollide(b)
	}

	if b.onCollide != nil {
		b.onCollide(a)
	}
}

// func combineCollisionIDs(id1, id2 collisionID) collisionID {
//if id1 < id2 {
//	return (id1 << 32) | id2
//	}
//	return (id2 << 32) | id1
//}
