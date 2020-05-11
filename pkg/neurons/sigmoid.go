package neurons

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

func NewSigmoidNeuron() *SigmoidNeuron {
	return &SigmoidNeuron{}
}

type SigmoidNeuron struct{}

func (neuron *SigmoidNeuron) Evaluate(input mat.Vector) mat.Vector {
	if input == nil {
		panic("neurons: sigmoid neuron got nil vector")
	}
	return neuron.apply(
		input,
		func(i int, value float64) float64 {
			return 1 / (1 + math.Exp(-value))
		},
	)
}

func (neuron *SigmoidNeuron) apply(vector mat.Vector, applier func(int, float64) float64) mat.Vector {
	length := vector.Len()
	output := mat.NewVecDense(length, nil)
	for i := 0; i < length; i++ {
		output.SetVec(i, applier(i, vector.AtVec(i)))
	}
	return output
}

func (neuron *SigmoidNeuron) Differentiate(output mat.Vector) mat.Vector {
	if output == nil {
		panic("neurons: sigmoid neuron got nil vector")
	}
	return neuron.apply(
		output,
		func(i int, value float64) float64 {
			return value * (1 - value)
		},
	)
}
