package initial_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/initial"
	"iio/pkg/test"
	"math/rand"
)

var _ = Describe("glorot", func() {
	initializer := initial.NewGlorotInitializer()
	BeforeEach(func() {
		rand.Seed(42)
	})
	Context("vector initialization", func() {
		test.With("should correctly initialize a vector of normally-distributed values", func() {
			test.Equate(
				initializer.InitializeVector(5),
				test.Vector(
					1.5536305584564762,
					0.12525608682704692,
					-0.4943748127704828,
					1.2440150150762053,
					0.1319784842710705,
				),
			)
		})
	})
	Context("matrix initialization", func() {
		test.With("should correctly initialize a matrix of normally-distributed values", func() {
			test.Equate(
				initializer.InitializeMatrix(5, 3),
				test.Matrix(
					5,
					3,
					3.1072611169129525,
					0.25051217365409384,
					-0.9887496255409656,
					2.4880300301524105,
					0.263956968542141,
					2.4127355794035426,
					-1.2515231701000755,
					1.2592199564734523,
					3.131040153048172,
					-1.668519954860832,
					-2.6367850294212674,
					1.7488798317284948,
					2.524323029527956,
					0.6911789875125134,
					-1.2980669082742273,
				),
			)
		})
	})
})
