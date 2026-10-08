package core

type Node2D struct {
	Node
	localTransform Transform
	worldTransform Transform
	isDirty        bool
}

func NewNode2D(name string) *Node2D {
	n := &Node2D{
		Node:           *NewNode(name),
		localTransform: NewTransform(ZeroVectore(), ZeroVectore(), 0),
		worldTransform: NewTransform(ZeroVectore(), ZeroVectore(), 0),
		isDirty:        true,
	}
	return n
}

// @overwite
func (node2d *Node2D) AddChild(child SceneNode) {
	node2d.Node.addChildWithParent(node2d, child)
}

func (node2d *Node2D) GetTransform() Transform {
	return node2d.localTransform
}

func (node2d *Node2D) SetTransform(transform Transform) {
	node2d.localTransform = transform
	node2d.MarkDirty()
}

func (node2d *Node2D) SetPosition(x, y float64) {
	node2d.localTransform.SetPosition(NewVector2D(x, y))
	node2d.MarkDirty()
}

func (node2d *Node2D) GetPosition() Vector2D {
	return node2d.localTransform.GetPosition()
}

func (node2D *Node2D) SetRotation(r float64) {
	node2D.localTransform.SetRotation(r)
	node2D.MarkDirty()
}

func (node2D *Node2D) GetRotation() float64 {
	return node2D.localTransform.GetRotation()
}

func (node2D *Node2D) SetScale(x, y float64) {
	node2D.localTransform.SetScale(x, y)
	node2D.MarkDirty()
}

func (node2D *Node2D) GetScale() Vector2D {
	return node2D.localTransform.GetScale()
}

func (node2D *Node2D) SetPivot(x, y float64) {
	node2D.localTransform.SetPivot(x, y)
	node2D.MarkDirty()
}

func (node2D *Node2D) GetPivot() Vector2D {
	return node2D.localTransform.GetPivot()
}

func (node2D *Node2D) MarkDirty() {
	if node2D.isDirty {
		return
	}

	node2D.isDirty = true
	node2D.Node.MarkDirty()
}

func (node2D *Node2D) GetWorldTransform() Transform {
	if !node2D.isDirty {
		return node2D.worldTransform
	}

	world := NewTransform(ZeroVectore(), ZeroVectore(), 0)
	if parent := node2D.GetParent(); parent != nil {
		if transfrom, ok := parent.(Transformable); ok {
			world = transfrom.GetWorldTransform()
		}
	}
	world.Concat(node2D.localTransform)
	node2D.worldTransform = world
	node2D.isDirty = false
	return world
}
