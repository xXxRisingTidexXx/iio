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
	rows, columns := nodes.Len(), activations.Len()
	weights := mat.NewDense(rows, columns, nil)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			weights.Set(i, j, nodes.AtVec(i)*activations.AtVec(j))
		}
	}
	return &Delta{weights, nodes}
}

type Delta struct {
	Weights mat.Matrix
	Biases  mat.Vector
}

func (delta *Delta) Equal(other *Delta) bool {
	return other != nil && mat.Equal(delta.Weights, other.Weights) && mat.Equal(delta.Biases, other.Biases)
}
