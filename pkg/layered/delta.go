package layered

import (
	"gonum.org/v1/gonum/mat"
)

func NewDelta(nodes, activations mat.Vector) *Delta {
	if nodes == nil {
		panic("layers: delta nodes can't be nil")
	}
	if activations == nil {
		panic("layers: delta activations can't be nil")
	}
	return &Delta{nodes, activations}
}

type Delta struct {
	Nodes       mat.Vector
	Activations mat.Vector
}

func (delta *Delta) Equal(other *Delta) bool {
	return other != nil &&
		mat.Equal(delta.Nodes, other.Nodes) &&
		mat.Equal(delta.Activations, other.Activations)
}
