package guts

import (
	"gonum.org/v1/gonum/mat"
)

type FeedForwardLayer struct {
	neuron  Neuron
	weights mat.Matrix
	biases  mat.Vector
}

func (layer *FeedForwardLayer) FeedForward(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedForwardLayer) ProduceNodes(diffs mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedForwardLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedForwardLayer) Update(delta *Delta) {
	panic("implement me")
}
