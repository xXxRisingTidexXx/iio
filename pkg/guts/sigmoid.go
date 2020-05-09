package guts

import (
	"gonum.org/v1/gonum/mat"
)

type SigmoidNeuron struct{}

func (neuron *SigmoidNeuron) Evaluate(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (neuron *SigmoidNeuron) Differentiate(activations mat.Vector) mat.Vector {
	panic("implement me")
}
