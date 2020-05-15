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
				test.Vector(0.0051188, 0.0146821, -0.000238, 1.200381),
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
		test.With("should correctly calculate diffs", func() {
			layer := layered.NewBasicLayer(
				neurons.NewSigmoidNeuron(),
				test.Matrix(
					4,
					3,
					-0.67474504,
					-1.21539734,
					0.3770968,
					-1.08962889,
					-1.69476444,
					1.75877852,
					-0.27093156,
					0.08103385,
					-1.70969782,
					0.27795207,
					3.22273382,
					0.98380333,
				),
				test.Vector(0.34674073, 1.13310711, -0.38119229, 0.992821716),
			)
			test.Equate(
				layer.BackPropagate(test.Vector(0.61270941, 0.31980839, 0.41249188, 0.00238712)),
				test.Vector(-0.8729886599576079, -1.2455664166060845, 0.09063487329007888),
			)
		})
	})
	Context("update", func() {
		test.With("should correctly apply layer updates", func() {
			layer := layered.NewBasicLayer(
				neurons.NewSigmoidNeuron(),
				test.Matrix(
					3,
					4,
					-0.31731328,
					0.24676347,
					-0.36625309,
					0.06265908,
					-0.58790104,
					-0.39845581,
					0.05905837,
					0.22153997,
					0.05141048,
					-0.03469448,
					-0.09033111,
					-0.4435566,
				),
				test.Vector(-0.21595326, -0.13819163, 0.31713667),
			)
			layer.Update(
				-0.01,
				&layered.Delta{
					Weights: test.Matrix(
						3,
						4,
						0.03436183,
						-0.17630402,
						0.0324084,
						-0.03850823,
						-0.0676922,
						0.06116763,
						0.10309995,
						0.09312801,
						-0.08392175,
						-0.03092124,
						0.03312634,
						0.09755451,
					),
					Biases: test.Vector(-0.04791742, -0.0185659, -0.1106335),
				},
			)
			test.Comply(
				layer,
				layered.NewBasicLayer(
					neurons.NewSigmoidNeuron(),
					test.Matrix(
						3,
						4,
						-0.3176568983,
						0.24852651020000002,
						-0.366577174,
						0.06304416230000001,
						-0.587224118,
						-0.3990674863,
						0.0580273705,
						0.22060868990000002,
						0.052249697500000004,
						-0.0343852676,
						-0.09066237340000001,
						-0.4445321451,
					),
					test.Vector(-0.21547408580000002, -0.138005971, 0.318243005),
				),
			)
		})
	})
})
