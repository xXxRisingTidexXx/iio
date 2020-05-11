package guts

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

const (
	sigmoidOne = 1.0
)

type SigmoidNeuron struct{}

// Calculates the elements of the incoming nodes array according
// to the sigmoid function (1/(1+e^-x))
func (neuron *SigmoidNeuron) Evaluate(activations mat.Vector) mat.Vector {
	vector := mat.NewVecDense(activations.Len(), nil)
	for i := 0; i < vector.Len(); i++ {
		vector.SetVec(i, sigmoidOne / (sigmoidOne + math.Pow(math.E, -activations.AtVec(i))))
	}
	return vector
}

// Calculates derivative of the elements of the incoming nodes array according
// to the sigmoid function (1/(1+e^-x)) -> (e^x/(1+e^x)^2)
func (neuron *SigmoidNeuron) Differentiate(activations mat.Vector) mat.Vector {
	vector := mat.NewVecDense(activations.Len(), nil)
	for i := 0; i < vector.Len(); i++ {
		vector.SetVec(i, math.Pow(math.E,
			activations.AtVec(i)) / math.Pow(sigmoidOne + math.Pow(math.E, activations.AtVec(i)), 2))
	}
	return vector
}
