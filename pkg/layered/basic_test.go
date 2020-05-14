package layered_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/layered"
	"iio/pkg/neurons"
	"iio/pkg/test"
)

var _ = Describe("basic", func() {
	Context("feed forward", func() {
		test.With("should correctly carry out direct propagation", func() {
			test.Comply(  // TODO: check a new neuron
				layered.NewBasicLayer(
					neurons.NewSigmoidNeuron(),
					test.Matrix(2, 1, 2, 0.3),
					test.Vector(3, 4),
				),
				layered.NewBasicLayer(
					neurons.NewSigmoidNeuron(),
					test.Matrix(2, 1, 2, 0.3),
					test.Vector(3, 4),
				),
			)
		})
	})
	Context("node production", func() {

	})
	Context("back propagation", func() {

	})
	Context("update", func() {

	})
})
