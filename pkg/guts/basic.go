package guts

import (
	"gonum.org/v1/gonum/mat"
)

func NewBasicLayer(neuron Neuron, weights *mat.Dense, biases *mat.VecDense) *BasicLayer {
	if neuron == nil || weights == nil || biases == nil {
		panic("guts: basic layer got nil argument(s)")
	}
	return &BasicLayer{neuron, weights, biases}
}

type BasicLayer struct {
	neuron  Neuron
	weights *mat.Dense
	biases  *mat.VecDense
}

func (layer *BasicLayer) FeedForward(activations mat.Vector) mat.Vector {
	if activations == nil {
		panic("guts: basic layer got nil vector")
	}
	rows, _ := layer.weights.Dims()
	input := mat.NewVecDense(rows, nil)
	input.MulVec(layer.weights, activations)
	input.AddVec(input, layer.biases)
	return layer.neuron.Evaluate(input)
}

func (layer *BasicLayer) ProduceNodes(diffs, activations mat.Vector) mat.Vector {
	if diffs == nil || activations == nil {
		panic("guts: basic layer got nil vector(s)")
	}
	nodes := mat.NewVecDense(activations.Len(), nil)
	nodes.MulElemVec(diffs, layer.neuron.Differentiate(activations))
	return nodes
}

func (layer *BasicLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	if nodes == nil {
		panic("guts: basic layer got nil vector")
	}
	_, columns := layer.weights.Dims()
	diffs := mat.NewVecDense(columns, nil)
	diffs.MulVec(layer.weights.T(), nodes)
	return diffs
}

func (layer *BasicLayer) Update(learningRate float64, delta *Delta) {
	if delta == nil {
		panic("guts: basic layer got nil delta")
	}
	layer.weights.Apply(
		func(i, j int, value float64) float64 {
			return value + learningRate*delta.Weights().At(i, j)
		},
		layer.weights,
	)
	layer.biases.AddScaledVec(layer.biases, learningRate, delta.Biases())
}
