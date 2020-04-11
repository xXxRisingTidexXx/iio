package networks

import (
	"iio/pkg/networks/layers"
	"iio/pkg/vectors"
)

// Primitive neural network implementation
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
