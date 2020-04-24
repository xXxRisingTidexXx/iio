package guts

import (
	"gonum.org/v1/gonum/mat"
)

type feedforwardLayer struct {
	neuron  Neuron
	weights mat.Matrix
	biases  mat.Vector
}

func (layer *feedforwardLayer) feedForward(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *feedforwardLayer) update(weightDeltas mat.Matrix, biasDeltas mat.Vector) {
	panic("implement me")
}
