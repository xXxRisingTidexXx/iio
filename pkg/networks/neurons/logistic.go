package neurons

import (
	"iio/pkg/vectors"
)

type LogisticNeuron struct {
	weights vectors.Vector
}

func (neuron *LogisticNeuron) Forward(activations vectors.Vector) float64 {
	panic("implement me")
}
