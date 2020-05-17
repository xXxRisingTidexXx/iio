package layered

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/neurons"
)

func NewBasicLayer(neuron neurons.Neuron, weights *mat.Dense, biases *mat.VecDense) Layer {
	if neuron == nil {
		panic("layers: basic layer neuron can't be nil")
	}
	if weights == nil {
		panic("layers: basic layer weights can't be nil")
	}
	if biases == nil {
		panic("layers: basic layer biases can't be nil")
	}
	rows, _ := weights.Dims()
	if length := biases.Len(); rows != length {
		panic(fmt.Sprintf("layers: basic layer matrix row numbers must equal, got %d & %d", rows, length))
	}
	return &basicLayer{neuron, weights, biases}
}

type basicLayer struct {
	neuron  neurons.Neuron
	weights *mat.Dense
	biases  *mat.VecDense
}

func (layer *basicLayer) Equal(other *basicLayer) bool {
	return other != nil &&
		cmp.Equal(layer.neuron, other.neuron) &&
		mat.Equal(layer.weights, other.weights) &&
		mat.Equal(layer.biases, other.biases)
}

func (layer *basicLayer) String() string {
	return fmt.Sprintf("{%s %v %v}", layer.neuron, layer.weights, layer.biases)
}

func (layer *basicLayer) FeedForward(activations mat.Vector) mat.Vector {
	if activations == nil {
		panic("layers: basic layer got nil activations")
	}
	rows, _ := layer.weights.Dims()
	input := mat.NewVecDense(rows, nil)
	input.MulVec(layer.weights, activations)
	input.AddVec(input, layer.biases)
	return layer.neuron.Evaluate(input)
}

func (layer *basicLayer) ProduceNodes(diffs, activations mat.Vector) mat.Vector {
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

func (layer *basicLayer) BackPropagate(nodes mat.Vector) mat.Vector {
	if nodes == nil {
		panic("layers: basic layer got nil nodes")
	}
	_, columns := layer.weights.Dims()
	diffs := mat.NewVecDense(columns, nil)
	diffs.MulVec(layer.weights.T(), nodes)
	return diffs
}

func (layer *basicLayer) Update(delta *Delta) {
	if delta == nil {
		panic("layers: basic layer got nil delta")
	}
	layer.weights.Add(layer.weights, delta.Weights)
	layer.biases.AddVec(layer.biases, delta.Biases)
}
