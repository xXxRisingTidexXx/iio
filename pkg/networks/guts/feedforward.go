package guts

import (
	"gonum.org/v1/gonum/mat"
)

type FeedforwardLayer struct {
	neuron  Neuron
	weights mat.Matrix
	biases  mat.Vector
}

func (layer *FeedforwardLayer) FeedForward(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedforwardLayer) ProduceNodes(diffs mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedforwardLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedforwardLayer) Update(delta *Delta) {
	panic("implement me")
}
