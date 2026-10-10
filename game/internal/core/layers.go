package core

import "github.com/hajimehoshi/ebiten/v2"

type Layers struct {
	layers []*Stack[func()]
}

func NewLayer() *Layers {
	layers := make([]*Stack[func()], 8)
	for i := range layers {
		layers[i] = NewStack[func()]()
	}
	return &Layers{
		layers: layers,
	}
}

func (l *Layers) getLayer(idx uint) *Stack[func()] {
	for uint(len(l.layers)) <= idx {
		l.layers = append(l.layers, NewStack[func()]())
	}
	return l.layers[idx]
}

func (l *Layers) AddDraw(layerIndex uint, f func()) {
	l.getLayer(layerIndex).Push(f)
}

func (l *Layers) DrawAll() {
	for _, stack := range l.layers {
		for !stack.IsEmpty() {
			if f, ok := stack.Pop(); ok {
				f()
			}
		}
	}
}

func (l *Layers) AddNode(layerIndex uint, node Drawable, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	opCopy := op
	l.AddDraw(layerIndex, func() {
		node.Draw(target, opCopy)
	})
}
