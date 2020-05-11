package init

import (
	"gonum.org/v1/gonum/mat"
)

func NewInitializer(kind Kind) Initializer {
	switch kind {
	case Zero:
		return &ZeroInitializer{}
	case Glorot:
		return &GlorotInitializer{}
	default:
		panic("initializers: undefined initializer kind")
	}
}

type Initializer interface {
	InitializeMatrix(int, int) *mat.Dense
	InitializeVector(int) *mat.VecDense
}
