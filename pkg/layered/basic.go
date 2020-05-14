package layered

import (
	"github.com/google/go-cmp/cmp"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/neurons"
)

func NewBasicLayer(neuron neurons.Neuron, weights *mat.Dense, biases *mat.VecDense) *BasicLayer {
	if neuron == nil {
		panic("layers: basic layer neuron can't be nil")
	}
	if weights == nil {
		panic("layers: basic layer weights can't be nil")
	}
	if biases == nil {
		panic("layers: basic layer biases can't be nil")
	}
	return &BasicLayer{neuron, weights, biases}
}

type BasicLayer struct {
	neuron  neurons.Neuron
	weights *mat.Dense
	biases  *mat.VecDense
}

func (layer *BasicLayer) Equal(other *BasicLayer) bool {
	return other != nil &&
		cmp.Equal(layer.neuron, other.neuron) &&
		mat.Equal(layer.weights, other.weights) &&
		mat.Equal(layer.biases, other.biases)
}

func (layer *BasicLayer) FeedForward(activations mat.Vector) mat.Vector {
	if activations == nil {
		panic("layers: basic layer got nil activations")
	}
	rows, _ := layer.weights.Dims()
	input := mat.NewVecDense(rows, nil)
	input.MulVec(layer.weights, activations)
	input.AddVec(input, layer.biases)
	return layer.neuron.Evaluate(input)
}

func (layer *BasicLayer) ProduceNodes(diffs, activations mat.Vector) mat.Vector {
	if diffs == nil {
		panic("layers: basic layer got nil diffs")
	}
	if activations == nil {
		panic("layers: basic layer got nil activations")
	}
	nodes := mat.NewVecDense(activations.Len(), nil)
	nodes.MulElemVec(diffs, layer.neuron.Differentiate(activations))
	return nodes
}

func (layer *BasicLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	if nodes == nil {
		panic("layers: basic layer got nil nodes")
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
			layer.weights.Set(i, j, layer.weights.At(i, j)+learningRate*delta.Weights.At(i, j))
		}
	}
	layer.biases.AddScaledVec(layer.biases, learningRate, delta.Biases)
}
