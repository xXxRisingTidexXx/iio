package guts

import (
	"gonum.org/v1/gonum/mat"
)

type OutputLayer struct {
	layer *feedforwardLayer
}

func (layer *OutputLayer) FeedForward(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *OutputLayer) BackPropagate(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *OutputLayer) Update(weightDeltas mat.Matrix, biasDeltas mat.Vector) {
	panic("implement me")
}
