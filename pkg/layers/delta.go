package layers

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewDelta(nodes, activations mat.Vector) *Delta {
	if nodes == nil || activations == nil {
		panic(fmt.Sprintf("layers: delta got nil vector"))
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
