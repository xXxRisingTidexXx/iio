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
	row, _ := layer.weights.Dims()
	z := mat.NewVecDense(row, nil)
	z.MulVec(layer.weights, activations)
	z.AddVec(z, layer.biases)
	return layer.neuron.Evaluate(z)
}

// Forms a node level
func (layer *FeedForwardLayer) ProduceNodes(diffs, activations mat.Vector) mat.Vector {
	resultDelta := mat.NewVecDense(diffs.Len(), nil)
	resultDelta.MulElemVec(diffs, layer.neuron.Differentiate(activations))
	return resultDelta
}

func (layer *FeedForwardLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	vector := mat.NewVecDense(nodes.Len(), nil)
	vector.MulVec(layer.weights.T(), nodes)
	return vector
}

func (layer *FeedForwardLayer) Update(delta *Delta) {
	panic("implement me")
}
