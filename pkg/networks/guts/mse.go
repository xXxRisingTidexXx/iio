package guts

import (
	"gonum.org/v1/gonum/mat"
)

type MSECostFunction struct{}

// Calculates errors
func (costFunction *MSECostFunction) Evaluate(activations mat.Vector, label int) mat.Vector {
	vector := mat.NewVecDense(activations.Len(), nil)
	vector.SetVec(label, 1)
	vector.SubVec(activations, vector)
	return vector
}
