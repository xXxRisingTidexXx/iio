package loss

import (
	"gonum.org/v1/gonum/mat"
)

type CostFunction interface {
	Evaluate(mat.Vector) int
	Differentiate(mat.Vector, int) mat.Vector
}
