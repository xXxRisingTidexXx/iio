package init

import (
	"gonum.org/v1/gonum/mat"
)

type Initializer interface {
	InitializeVector(int) *mat.VecDense
	InitializeMatrix(int, int) *mat.Dense
}
