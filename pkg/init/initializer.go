package init

import (
	"gonum.org/v1/gonum/mat"
)

type Initializer interface {
	InitializeMatrix(int, int) *mat.Dense
	InitializeVector(int) *mat.VecDense
}
