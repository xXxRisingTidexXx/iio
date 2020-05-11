package layers

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/neurons"
)

func NewBasicLayer(kind neurons.Kind, weights *mat.Dense, biases *mat.VecDense) *BasicLayer {
	if weights == nil || biases == nil {
		panic("layers: basic layer got nil vector(s)")
	}
	return &BasicLayer{neurons.NewNeuron(kind), weights, biases}
}

type BasicLayer struct {
	neuron  neurons.Neuron
	weights *mat.Dense
	biases  *mat.VecDense
}

func (layer *BasicLayer) FeedForward(activations mat.Vector) mat.Vector {
	if activations == nil {
		panic("layers: basic layer got nil vector")
	}
	rows, _ := layer.weights.Dims()
	input := mat.NewVecDense(rows, nil)
	input.MulVec(layer.weights, activations)
	input.AddVec(input, layer.biases)
	return layer.neuron.Evaluate(input)
}

func (layer *BasicLayer) ProduceNodes(diffs, activations mat.Vector) mat.Vector {
	if diffs == nil || activations == nil {
		panic("layers: basic layer got nil vector(s)")
	}
	nodes := mat.NewVecDense(activations.Len(), nil)
	nodes.MulElemVec(diffs, layer.neuron.Differentiate(activations))
	return nodes
}

func (layer *BasicLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	if nodes == nil {
		panic("layers: basic layer got nil vector")
	}
	_, columns := layer.weights.Dims()
	diffs := mat.NewVecDense(columns, nil)
	diffs.MulVec(layer.weights.T(), nodes)
	return diffs
}

func (layer *BasicLayer) Update(learningRate float64, delta *Delta) {
	if delta == nil {
		panic("layers: basic layer got nil delta")
	}
	rows, columns := layer.weights.Dims()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			layer.weights.Set(i, j, layer.weights.At(i, j)+learningRate*delta.Weights().At(i, j))
		}
	}
	layer.biases.AddScaledVec(layer.biases, learningRate, delta.Biases())
}
