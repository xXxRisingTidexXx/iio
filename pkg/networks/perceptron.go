package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
)

type Perceptron struct{
	network *feedforwardNetwork
}

func (perceptron *Perceptron) Train(samples *sampling.Samples) {
	panic("implement me")
}

func (perceptron *Perceptron) Validate(samples *sampling.Samples) {
	panic("implement me")
}

func (perceptron *Perceptron) Test(samples *sampling.Samples) Report {
	panic("implement me")
}

func (perceptron *Perceptron) Evaluate(activations mat.Vector) int {
	panic("implement me")
}
