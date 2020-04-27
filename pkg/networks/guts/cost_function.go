package guts

import (
	"gonum.org/v1/gonum/mat"
)

type CostFunction interface {
	Evaluate(mat.Vector, int) mat.Vector
}
