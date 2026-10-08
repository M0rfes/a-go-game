package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Node2D
	texture *ebiten.Image
	layer   int
	visible bool
}

func NewSprite(name string, texture *ebiten.Image, layer int, centerPivot bool) *Sprite {
	s := &Sprite{
		Node2D:  *NewNode2D(name),
		texture: texture,
		layer:   layer,
		visible: true,
	}

	if texture != nil && centerPivot {
		s.SetPivotToCenter()
	}

	return s
}

func (s *Sprite) SetPivotToCenter() {
	w := float64(s.texture.Bounds().Dx() / 2)
	h := float64(s.texture.Bounds().Dy() / 2)
	s.SetPivot(w, h)
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}

func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

func (s *Sprite) GetLayer() int {
	return s.layer
}

func (s *Sprite) SetLayer(layer int) {
	s.layer = layer
}

func (s *Sprite) GetVicible() bool {
	return s.visible
}

func (s *Sprite) SetVisible(visible bool) {
	s.visible = visible
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if s.texture == nil || !s.visible {
		return
	}
	target.DrawImage(s.texture, op)
}
