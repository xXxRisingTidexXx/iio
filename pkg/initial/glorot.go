package initial

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

func NewGlorotInitializer() *GlorotInitializer {
	return &GlorotInitializer{}
}

// Glorot (Xavier) normal sigmoid initializer.
type GlorotInitializer struct{}

func (initializer *GlorotInitializer) String() string {
	return "glorot"
}

func (initializer *GlorotInitializer) InitializeVector(length int) *mat.VecDense {
	vector := mat.NewVecDense(length, nil)
	for i := 0; i < length; i++ {
		vector.SetVec(i, rand.NormFloat64())
	}
	return vector
}

func (initializer *GlorotInitializer) InitializeMatrix(rows, columns int) *mat.Dense {
	sigma := math.Sqrt(32 / float64(rows+columns))
	matrix := mat.NewDense(rows, columns, nil)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matrix.Set(i, j, rand.NormFloat64()*sigma)
		}
	}
	return matrix
}
