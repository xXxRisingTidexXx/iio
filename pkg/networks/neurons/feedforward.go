package neurons

import (
	"iio/pkg/vectors"
)

type feedforwardNeuron struct {
	weights vectors.Vector
	bias    float64
}

func (neuron *feedforwardNeuron) feedForward(activations vectors.Vector) float64 {
	panic("implement me")
}

func (neuron *feedforwardNeuron) backPropagate(delta float64) vectors.Vector {
	panic("implement me")
}

func (neuron *feedforwardNeuron) update(weightDeltas vectors.Vector, biasDelta float64) {
	panic("implement me")
}
