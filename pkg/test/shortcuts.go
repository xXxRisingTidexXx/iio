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

func Matrix(rows, columns int, values ...float64) *mat.Dense {
	if rows <= 0 || columns <= 0 {
		panic("test: matrix dimensions can't be non-positive")
	}
	length := len(values)
	if values == nil || length == 0 {
		panic("test: matrix can't be empty")
	}
	if rows*columns != length {
		panic("test: matrix dimension inconsistency")
	}
	return mat.NewDense(rows, columns, values)
}

func Equate(a, b mat.Matrix) {
	gomega.Expect(mat.Equal(a, b)).To(gomega.BeTrue())
}
