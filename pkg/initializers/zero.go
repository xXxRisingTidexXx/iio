package initializers

import (
	"gonum.org/v1/gonum/mat"
)

type ZeroInitializer struct{}

func (initializer *ZeroInitializer) InitializeMatrix(rows, columns int) *mat.Dense {
	return mat.NewDense(rows, columns, nil)
}

func (initializer *ZeroInitializer) InitializeVector(length int) *mat.VecDense {
	return mat.NewVecDense(length, nil)
}
