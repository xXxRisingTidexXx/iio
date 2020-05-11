package guts

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewDelta(nodes, activations mat.Vector) *Delta {
	if nodes == nil || activations == nil {
		panic(fmt.Sprintf("guts: delta got nil vector"))
	}
	weights := mat.NewDense(nodes.Len(), activations.Len(), nil)
	weights.Apply(
		func(i, j int, value float64) float64 {
			return nodes.AtVec(i) * activations.AtVec(j)
		},
		weights,
	)
	return &Delta{weights, nodes}
}

type Delta struct {
	weights mat.Matrix
	biases  mat.Vector
}

func (delta *Delta) Weights() mat.Matrix {
	return delta.weights
}

func (delta *Delta) Biases() mat.Vector {
	return delta.biases
}
