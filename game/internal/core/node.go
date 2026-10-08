package core

import "sync/atomic"

var globalNodeID uint64

type Node struct {
	id       uint64
	name     string
	children []SceneNode
	parent   SceneNode
}

func NewNode(name string) *Node {
	id := atomic.AddUint64(&globalNodeID, 1)
	return &Node{id: id, name: name, children: make([]SceneNode, 0)}
}

func (n *Node) GetID() uint64 {
	return n.id
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) GetParent() SceneNode {
	return n.parent
}

// GetChildren implements [SceneNode].
func (n *Node) GetChildren() []SceneNode {
	return n.children
}

func (n *Node) addChildWithParent(logicalParent SceneNode, child SceneNode) {
	child.AddParent(logicalParent)
	n.children = append(n.children, child)
}

func (n *Node) AddChild(child SceneNode) {
	n.addChildWithParent(n, child)
}

// AddParent implements [SceneNode].
func (n *Node) AddParent(parent SceneNode) {
	n.parent = parent
}

// DetachChild implements [SceneNode].
func (n *Node) DetachChild(node SceneNode) bool {
	for i, c := range n.children {
		if c == node {
			n.children[i] = n.children[len(n.children)-1]
			n.children = n.children[:len(n.children)-1]
			node.AddParent(nil)
			return true
		}
	}
	return false
}

// MarkDirty implements [SceneNode].
func (n *Node) MarkDirty() {
	for _, n := range n.children {
		n.MarkDirty()
	}
}
