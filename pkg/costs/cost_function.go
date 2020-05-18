package costs

import (
	"gonum.org/v1/gonum/mat"
)

type CostFunction interface {
	Evaluate(mat.Vector) int
	Cost(mat.Vector, int) float64
	Differentiate(mat.Vector, int) mat.Vector
}
