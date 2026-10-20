package core

type CollisionCircul struct {
	radius float64
}

func NewCollisionCircle(radius float64) *CollisionCircul {
	return &CollisionCircul{radius: radius}
}

func (c *CollisionCircul) GetRadius() float64 {
	return c.radius
}

func circlesOverlap(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x2 - x1
	dy := y2 - y1

	dist := dx*dx + dy*dy

	sum := r1 + r2
	return dist < sum*sum
}
