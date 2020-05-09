package guts

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

type SigmoidNeuron struct{}

func (neuron *SigmoidNeuron) Evaluate(input mat.Vector) mat.Vector {
	if input == nil {
		panic("guts: sigmoid got nil vector")
	}
	return neuron.apply(
		input,
		func(i int, value float64) float64 {
			return 1 / (1 + math.Exp(-value))
		},
	)
}

func (neuron *SigmoidNeuron) apply(input mat.Vector, applier func(int, float64) float64) mat.Vector {
	length := input.Len()
	output := mat.NewVecDense(length, nil)
	for i := 0; i < length; i++ {
		output.SetVec(i, applier(i, input.AtVec(i)))
	}
	return output
}

func (neuron *SigmoidNeuron) Differentiate(input mat.Vector) mat.Vector {
	if input == nil {
		panic("guts: sigmoid got nil vector")
	}
	return neuron.apply(
		input,
		func(i int, value float64) float64 {
			return value * (1 - value)
		},
	)
}
