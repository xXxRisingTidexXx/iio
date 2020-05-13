package layered_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/layered"
	"iio/pkg/test"
)

var _ = Describe("delta", func() {
	test.With("should correctly create new delta from nodes and activations", func() {
		test.Comply(
			layered.NewDelta(
				test.Vector(0.4502, 0.0000281, 0.92837),
				test.Vector(0.30113, 0.000007, 0.01198, 0.28767, 0.88891),
			),
			&layered.Delta{
				Weights: test.Matrix(
					3,
					5,
					0.135568726,
					3.1514e-6,
					0.005393396,
					0.129509034,
					0.400187282,
					8.461753e-6,
					1.9669999999999998e-10,
					3.3663799999999997e-7,
					8.083526999999999e-6,
					2.4978371e-5,
					0.2795600581,
					6.49859e-6,
					0.011121872599999999,
					0.26706419789999997,
					0.8252373767,
				),
				Biases: test.Vector(0.4502, 0.0000281, 0.92837),
			},
		)
	})
})
