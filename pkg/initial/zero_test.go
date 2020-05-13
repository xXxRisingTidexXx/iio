package initial_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/initial"
	"iio/pkg/test"
)

var _ = Describe("zero", func() {
	initializer := initial.NewZeroInitializer()
	Context("vector initialization", func() {
		test.With("should correctly initialize a vector of zeros", func() {
			test.Equate(initializer.InitializeVector(5), test.Vector(0, 0, 0, 0, 0))
		})
	})
	Context("matrix initialization", func() {
		test.With("should correctly initialize a matrix of zeros", func() {
			test.Equate(initializer.InitializeMatrix(3, 3), test.Matrix(3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0))
		})
	})
})
