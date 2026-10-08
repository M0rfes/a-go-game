package core

type SceneNode interface {
	GetID() uint64
	GetName() string
	GetChildren() []SceneNode
	GetParent() SceneNode

	AddChild(child SceneNode)
	AddParent(parent SceneNode)

	DetachChild(node SceneNode) bool
	MarkDirty()
}
