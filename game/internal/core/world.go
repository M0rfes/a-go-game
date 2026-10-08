package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	rootScene SceneNode
}

func NewWorld() *World {
	return &World{rootScene: NewNode("root")}
}

func (world *World) AddNodeToDefaultLayer(node SceneNode) {
	world.rootScene.AddChild(node)
}

func (world *World) Update() {
	world.updateNode(world.rootScene)
}

func (world *World) updateNode(node SceneNode) {
	if node == nil {
		return
	}
	for _, child := range node.GetChildren() {
		world.updateNode(child)
	}
}

func (world *World) Draw(target *ebiten.Image) {
	world.drawNodes(world.rootScene, ebiten.GeoM{}, target)
}

func (world *World) drawNodes(node SceneNode, parentGeoM ebiten.GeoM, target *ebiten.Image) {
	if node == nil {
		return
	}

	if drawable, ok := node.(Drawable); ok {
		op := &ebiten.DrawImageOptions{}
		op.GeoM = buildGeoMFromTransform(drawable.GetWorldTransform())
		drawable.Draw(target, op)
	}
	for _, child := range node.GetChildren() {
		world.drawNodes(child, parentGeoM, target)
	}
}

func (world *World) RemoveNode(node SceneNode) bool {
	parent := node.GetParent()
	if parent == nil {
		return false
	}

	if !parent.DetachChild(node) {
		return false
	}

	node.AddParent(nil)
	return true
}

func buildGeoMFromTransform(t Transform) ebiten.GeoM {
	g := ebiten.GeoM{}
	pivot := t.GetPivot()
	position := t.GetPosition()
	scale := t.GetScale()
	rotation := t.GetRotation()

	g.Translate(-pivot.X(), -pivot.Y())
	g.Scale(scale.X(), scale.Y())
	g.Rotate(rotation)
	g.Translate(position.X(), position.Y())

	return g
}
