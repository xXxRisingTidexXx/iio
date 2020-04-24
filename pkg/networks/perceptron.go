package networks

import (
	"iio/pkg/sampling"
	"iio/pkg/vectors"
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

func (perceptron *Perceptron) Evaluate(activations vectors.Vector) byte {
	panic("implement me")
}
