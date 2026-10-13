package core

import "github.com/hajimehoshi/ebiten/v2"

type Camera struct {
	*Node2D
	width      uint
	height     uint
	surface    *ebiten.Image
	nodeTollow Transformable
}

func NewCamera(width, height uint) *Camera {
	c := &Camera{
		Node2D: NewNode2D("camera"),
		width:  width,
		height: height,
	}
	c.surface = ebiten.NewImage(int(width), int(height))
	return c
}

func (c *Camera) SetFollow(node Transformable) {
	c.nodeTollow = node
}

func (c *Camera) Update() {
	if c.nodeTollow == nil {
		return
	}

	wt := c.nodeTollow.GetWorldTransform()
	px := wt.GetPosition().X()
	py := wt.GetPosition().Y()
	c.SetPosition(px-float64(c.width)/2, py-float64(c.height)/2)
}

func (c *Camera) ApplyOffset(op *ebiten.DrawImageOptions) {
	px := c.GetPosition().X()
	py := c.GetPosition().Y()
	op.GeoM.Translate(-px, -py)
}

func (c *Camera) DrawToScreen(target *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	target.DrawImage(c.surface, op)
}

func (c *Camera) GetSurface() *ebiten.Image {
	return c.surface
}
