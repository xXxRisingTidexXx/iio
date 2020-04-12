package neurons

import (
	"iio/pkg/vectors"
)

// Casual neuron's implementation, which leverages logistic
// function (https://en.wikipedia.org/wiki/Logistic_function)
// to finalize the output.
type LogisticNeuron struct {
	weights vectors.Vector
}

func (neuron *LogisticNeuron) Forward(activations vectors.Vector) float64 {
	panic("implement me")
}
