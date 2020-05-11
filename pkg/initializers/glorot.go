package initializers

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

type GlorotInitializer struct{}

func (initializer *GlorotInitializer) InitializeMatrix(rows, columns int) *mat.Dense {
	sigma := math.Sqrt(1 / float64(columns))
	matrix := mat.NewDense(rows, columns, nil)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matrix.Set(i, j, rand.NormFloat64()*sigma)
		}
	}
	return matrix
}

func (initializer *GlorotInitializer) InitializeVector(length int) *mat.VecDense {
	vector := mat.NewVecDense(length, nil)
	for i := 0; i < length; i++ {
		vector.SetVec(i, rand.NormFloat64())
	}
	return vector
}
