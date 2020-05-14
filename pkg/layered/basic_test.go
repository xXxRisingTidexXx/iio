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
			layer := layered.NewBasicLayer(
				neurons.NewSigmoidNeuron(),
				test.Matrix(2, 4, 2.001, 0.3, -3.501, 0, 0.00193, -0.0038, -1.98163, 1.39382),
				test.Vector(0.3, 0.00054),
			)
			test.Equate(
				layer.FeedForward(test.Vector(0.001398, 0.002438, 1.345e-5, 0.659981)),
				test.Vector(0.5752934282450957, 0.7151239066703989),
			)
		})
	})
	Context("node production", func() {
		test.With("should correctly produce node deltas", func() {
			layer := layered.NewBasicLayer(
				neurons.NewSigmoidNeuron(),
				test.Matrix(4, 2, 0.02348, 0.3891, -3.501, 0.234, 0.00193, -0.3, -2, 1.0022),
				test.Vector(0.0051, 0.01),
			)
			test.Equate(
				layer.ProduceNodes(
					test.Vector(0.000238, 0.921, 0.572, 0.133713),
					test.Vector(0.11239, 0.4461, 0.81237, 0.00238),
				),
				test.Vector(2.37425201202e-5, 0.22757430159000006, 0.08718709033319998, 0.0003174795360828),
			)
		})
	})
	Context("back propagation", func() {

	})
	Context("update", func() {

	})
})
