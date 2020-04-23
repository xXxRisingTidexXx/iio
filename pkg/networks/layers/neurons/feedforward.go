package neurons

import (
	"iio/pkg/vectors"
)

type feedforwardNeuron struct {
	bunch *Bunch
}

func (neuron *feedforwardNeuron) feedForward(activations vectors.Vector) float64 {
	panic("implement me")
}

func (neuron *feedforwardNeuron) backPropagate(delta float64) vectors.Vector {
	panic("implement me")
}

func (neuron *feedforwardNeuron) update(bunch *Bunch) {
	panic("implement me")
}
