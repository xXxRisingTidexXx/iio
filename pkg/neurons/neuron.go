package neurons

import (
	"gonum.org/v1/gonum/mat"
)

func NewNeuron(kind Kind) Neuron {
	switch kind {
	case Input:
		panic("neurons: can't instantiate input neuron")
	case Sigmoid:
		return &SigmoidNeuron{}
	default:
		panic("neurons: undefined neuron kind")
	}
}

type Neuron interface {
	Evaluate(mat.Vector) mat.Vector
	Differentiate(mat.Vector) mat.Vector
}
