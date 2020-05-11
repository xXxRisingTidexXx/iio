package costs

import (
	"gonum.org/v1/gonum/mat"
)

func NewCostFunction(kind Kind) CostFunction {
	switch kind {
	case MSE:
		return &MSECostFunction{}
	default:
		panic("loss: undefined cost function kind")
	}
}

type CostFunction interface {
	Evaluate(mat.Vector) int
	Differentiate(mat.Vector, int) mat.Vector
}
