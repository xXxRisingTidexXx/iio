package guts

import (
	"gonum.org/v1/gonum/mat"
)

type MSECostFunction struct {}

func (costFunction *MSECostFunction) Evaluate(activations mat.Vector, label int) mat.Vector {
	panic("implement me")
}

