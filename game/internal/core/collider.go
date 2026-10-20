package core

type Collider struct {
	*Node2D
	shape     *CollisionCircul
	mask      CollisionMask
	onCollide func(other *Collider)
}

func (c *Collider) SetOnCollide(fn func(other *Collider)) {
	c.onCollide = fn
}

func (c *Collider) GetShape() *CollisionCircul {
	return c.shape
}

func (c *Collider) GetMask() CollisionMask {
	return c.mask
}

func (c *Collider) CanCollideWith(other *Collider) bool {
	return c.mask.CanCollideWith(other.mask)
}
