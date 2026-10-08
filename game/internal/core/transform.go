package core

type Transform struct {
	position Vector2D
	pivot    Vector2D
	scale    Vector2D
	rotation float64
}

func NewTransform(position, pivot Vector2D, rotation float64) Transform {
	return Transform{
		position: position,
		pivot:    pivot,
		scale:    NewVector2D(1, 1),
		rotation: rotation,
	}
}

func (t *Transform) GetPosition() Vector2D {
	return t.position
}
func (t *Transform) SetPosition(v Vector2D) {
	t.position.SetPosition(v)
}
func (t *Transform) GetPivot() Vector2D {
	return t.pivot
}
func (t *Transform) GetRotation() float64 {
	return t.rotation
}
func (t *Transform) GetScale() Vector2D {
	return t.scale
}

func (t *Transform) SetPivot(x, y float64) {
	t.pivot.SetPosition(NewVector2D(x, y))
}
func (t *Transform) SetRotation(r float64) {
	t.rotation = r
}
func (t *Transform) SetScale(x, y float64) {
	t.scale.SetPosition(NewVector2D(x, y))
}

func (t *Transform) Translate(x, y float64) {
	t.position.x += x
	t.position.y += y
}
func (t *Transform) Rotate(r float64) {
	t.rotation += r
}

func (t *Transform) Concat(other Transform) {
	sx := other.position.X() * t.scale.X()
	sy := other.position.Y() * t.scale.Y()
	rotation := NewVector2D(sx, sy).Rotate(t.rotation)

	t.Translate(rotation.x, rotation.y)
	t.scale.SetPosition(NewVector2D(t.scale.X()*other.scale.X(), t.scale.Y()*other.scale.Y()))
	t.Rotate(other.rotation)
	t.pivot.SetPosition(other.GetPivot())
}
