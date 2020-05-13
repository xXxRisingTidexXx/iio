package init

import (
	"gonum.org/v1/gonum/mat"
)

func NewZeroInitializer() *ZeroInitializer {
	return &ZeroInitializer{}
}

type ZeroInitializer struct{}

func (initializer *ZeroInitializer) InitializeVector(length int) *mat.VecDense {
	return mat.NewVecDense(length, nil)
}

func (initializer *ZeroInitializer) InitializeMatrix(rows, columns int) *mat.Dense {
	return mat.NewDense(rows, columns, nil)
}
