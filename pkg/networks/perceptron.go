package networks

import (
	"iio/pkg/networks/layers"
	"iio/pkg/vectors"
)

// Primitive neural network implementation, which leverages
// multi-layer architecture. It uses sequentially-ordered hidden
// layers transmitting the activation signal from neurons to
// neurons towards the output layer. Each neuron of a following
// layer is connected to each neuron of the previous one.
type Perceptron struct {
	layers []layers.Layer
}

func (perceptron *Perceptron) Train(samples []*Sample) float64 {
	panic("implement me")
}

func (perceptron *Perceptron) Test(samples []*Sample) float64 {
	panic("implement me")
}

func (perceptron *Perceptron) Classify(vector vectors.Vector) byte {
	panic("implement me")
}

func (perceptron *Perceptron) Dump(path string) error {
	panic("implement me")
}
