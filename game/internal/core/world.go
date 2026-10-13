package core

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

const DefaultLayerIndex = 0

type World struct {
	rootScene  SceneNode
	layerRoots []SceneNode
	drawLayers *Layers
	drawTarget *ebiten.Image

	camera *Camera
}

func NewWorld(w, h uint) *World {
	return &World{rootScene: NewNode("root"), layerRoots: make([]SceneNode, 0), drawLayers: NewLayer(), camera: NewCamera(w, h)}
}

func (world *World) AddNodeToDefaultLayer(node SceneNode) {
	world.AddNodeToLayer(node, DefaultLayerIndex)
}

func (world *World) Update() {
	world.camera.Update()
	world.updateNode(world.rootScene)
}

func (world *World) updateNode(node SceneNode) {
	if node == nil {
		return
	}
	if u, ok := node.(Updatable); ok {
		u.Update()
	}
	for _, child := range node.GetChildren() {
		world.updateNode(child)
	}
}

func (w *World) prepareDrawSurface() {
	w.drawTarget = w.camera.GetSurface()
	w.drawTarget.Fill(color.RGBA{40, 40, 50, 255})
}

func (world *World) Draw(target *ebiten.Image) {
	world.prepareDrawSurface()

	for i := range world.layerRoots {
		world.drawNodes(world.layerRoots[i], uint(i))
	}
	world.drawLayers.DrawAll()
	world.camera.DrawToScreen(target)
}

func (world *World) drawNodes(node SceneNode, layerIndex uint) {
	if node == nil {
		return
	}
	childrens := node.GetChildren()
	sort.Slice(childrens, func(i, j int) bool {
		li, lj := 0, 0
		if d, ok := childrens[i].(Drawable); ok {
			li = int(d.GetLayer())
		}

		if d, ok := childrens[j].(Drawable); ok {
			lj = int(d.GetLayer())
		}
		return li > lj
	})

	for _, child := range childrens {
		world.drawNodes(child, layerIndex)
	}

	if d, ok := node.(Drawable); ok {
		op := &ebiten.DrawImageOptions{}
		if t, ok := node.(Transformable); ok {
			op.GeoM = buildGeoMFromTransform(t.GetWorldTransform())
		}
		world.camera.ApplyOffset(op)
		world.drawLayers.AddNode(layerIndex, d, world.drawTarget, op)
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

func (world *World) getLayer(layerIndex uint) SceneNode {
	for layerIndex >= uint(len(world.layerRoots)) {
		root := NewNode("layer_root")
		world.layerRoots = append(world.layerRoots, root)
		world.rootScene.AddChild(root)
	}
	return world.layerRoots[layerIndex]
}

func (world *World) AddNodeToLayer(node SceneNode, layerIndex uint) {
	world.getLayer(layerIndex).AddChild(node)
}

func (world *World) Camera() *Camera {
	return world.camera
}
