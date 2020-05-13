package test

import (
	"github.com/onsi/gomega"
	"gonum.org/v1/gonum/mat"
)

func Vector(values ...float64) *mat.VecDense {
	length := len(values)
	if values == nil || length == 0 {
		panic("test: vector can't be empty")
	}
	return mat.NewVecDense(length, values)
}

func Zeros(length int) *mat.VecDense {
	if length <= 0 {
		panic("test: vector of zeros can't be empty")
	}
	return mat.NewVecDense(length, nil)
}

func Equate(a, b mat.Matrix) {
	gomega.Expect(mat.Equal(a, b)).To(gomega.BeTrue())
}
