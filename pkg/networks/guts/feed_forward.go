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

// Forms a node level
func (layer *FeedForwardLayer) ProduceNodes(diffs mat.Vector) mat.Vector {
	row, _ := layer.weights.Dims()
	vector := mat.NewVecDense(row, nil)
	vector.MulVec(layer.weights, diffs)
	vector.AddVec(vector, layer.biases)
	return layer.neuron.Evaluate(vector)
}

func (layer *FeedForwardLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *FeedForwardLayer) Update(delta *Delta) {
	panic("implement me")
}
