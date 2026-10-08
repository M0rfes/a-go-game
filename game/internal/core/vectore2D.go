package core

import "math"

type Vector2D struct {
	x, y float64
}

func NewVector2D(x, y float64) Vector2D {
	return Vector2D{x, y}
}

func ZeroVectore() Vector2D {
	return Vector2D{0, 0}
}

func (vec *Vector2D) X() float64 {
	return vec.x
}

func (vec *Vector2D) Y() float64 {
	return vec.y
}

func (vec *Vector2D) SetX(x float64) {
	vec.x = x
}

func (vec *Vector2D) SetY(y float64) {
	vec.y = y
}

func (vec *Vector2D) SetPosition(other Vector2D) {
	vec.x, vec.y = other.x, other.y
}

func (vec Vector2D) Rotate(radians float64) Vector2D {
	s, c := math.Sin(radians), math.Cos(radians)
	// [vec.x] [cosø,  sinø]
	// [vec.y] [-sinø, cosø]
	return Vector2D{vec.x*c - vec.y*s, vec.x*s + vec.y*c}
}
