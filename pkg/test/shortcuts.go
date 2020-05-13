package test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"gonum.org/v1/gonum/mat"
)

func With(text string, body func()) {
	ginkgo.It(text, func() {
		defer func() {
			gomega.Expect(recover()).To(gomega.BeNil())
		}()
		body()
	})
}

func Spare(text string, body func()) {
	ginkgo.It(text, func() {
		defer func() {
			gomega.Expect(recover()).NotTo(gomega.BeNil())
		}()
		body()
	})
}

func Equate(a, b mat.Matrix) {
	gomega.Expect(mat.Equal(a, b)).To(gomega.BeTrue())
}

func Comply(a, b interface{}) {
	gomega.Expect(cmp.Equal(a, b)).To(gomega.BeTrue())
}


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
