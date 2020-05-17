package initial

import (
	"gonum.org/v1/gonum/mat"
)

func NewZeroInitializer() Initializer {
	return &zeroInitializer{}
}

type zeroInitializer struct{}

func (initializer *zeroInitializer) String() string {
	return "zero"
}

func (initializer *zeroInitializer) InitializeVector(length int) *mat.VecDense {
	return mat.NewVecDense(length, nil)
}

func (initializer *zeroInitializer) InitializeMatrix(rows, columns int) *mat.Dense {
	return mat.NewDense(rows, columns, nil)
}
