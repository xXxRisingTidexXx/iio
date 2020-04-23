package neurons

import (
	"iio/pkg/vectors"
)

type SigmoidNeuron struct{
	neuron *feedforwardNeuron
}

func (neuron *SigmoidNeuron) FeedForward(activations vectors.Vector) float64 {
	panic("implement me")
}

func (neuron *SigmoidNeuron) Derivative(activation float64) float64 {
	panic("implement me")
}

func (neuron *SigmoidNeuron) BackPropagate(delta float64) vectors.Vector {
	panic("implement me")
}

func (neuron *SigmoidNeuron) Update(weightDeltas vectors.Vector, biasDelta float64) {
	panic("implement me")
}
