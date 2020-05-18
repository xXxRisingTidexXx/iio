package initial

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

func NewGlorotInitializer() Initializer {
	return &glorotInitializer{}
}

// Glorot (Xavier) normal sigmoid initializer.
type glorotInitializer struct{}

func (initializer *glorotInitializer) String() string {
	return "glorot"
}

func (initializer *glorotInitializer) InitializeVector(length int) *mat.VecDense {
	vector := mat.NewVecDense(length, nil)
	for i := 0; i < length; i++ {
		vector.SetVec(i, rand.NormFloat64())
	}
	return vector
}

func (initializer *glorotInitializer) InitializeMatrix(rows, columns int) *mat.Dense {
	sigma := math.Sqrt(32 / float64(rows+columns))
	matrix := mat.NewDense(rows, columns, nil)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matrix.Set(i, j, rand.NormFloat64()*sigma)
		}
	}
	return matrix
}
