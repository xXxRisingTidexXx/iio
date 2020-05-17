package layered

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/neurons"
)

func NewOptions(neuron neurons.Neuron, weights *mat.Dense, biases *mat.VecDense) *Options {
	if neuron == nil {
		panic("layers: layer neuron can't be nil")
	}
	if weights == nil {
		panic("layers: layer weights can't be nil")
	}
	if biases == nil {
		panic("layers: layer biases can't be nil")
	}
	rows, _ := weights.Dims()
	if length := biases.Len(); rows != length {
		panic(fmt.Sprintf("layers: layer matrix row numbers must equal, got %d & %d", rows, length))
	}
	return &Options{neuron, weights, biases}
}

type Options struct {
	Neuron  neurons.Neuron
	Weights *mat.Dense
	Biases  *mat.VecDense
}
