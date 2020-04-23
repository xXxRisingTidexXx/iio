package networks

import (
	"iio/pkg/networks/reports"
	"iio/pkg/sampling"
	"iio/pkg/vectors"
)

// Primitive neural network implementation, which leverages
// multi-layer architecture. It uses sequentially-ordered hidden
// layers transmitting the activation signal from neurons to
// neurons towards the output layer. Each neuron of a following
// layer is connected to each neuron of the previous one.
type Perceptron struct {}

func (perceptron *Perceptron) Train(samples []*sampling.Sample) {
	panic("implement me")
}

func (perceptron *Perceptron) Validate(samples []*sampling.Sample) {
	panic("implement me")
}

func (perceptron *Perceptron) Test(samples []*sampling.Sample) reports.Report {
	panic("implement me")
}

func (perceptron *Perceptron) Evaluate(activations vectors.Vector) byte {
	panic("implement me")
}
