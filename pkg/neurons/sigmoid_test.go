package neurons_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/neurons"
	"iio/pkg/test"
)

var _ = Describe("sigmoid", func() {
	neuron := neurons.NewSigmoidNeuron()
	Context("evaluation", func() {
		It("should correctly apply logistic function to vector", func() {
			test.Equate(
				neuron.Evaluate(
					test.Vector(-0.23, 0.7682, 2.92881, 0.00271, 5.782, 0, -2391, 1037.5, -5.820, -30.00041),
				),
				test.Vector(
					0.44275214540144436,
					0.6831313890266155,
					0.9492523806705182,
					0.5006774995853647,
					0.996926928727142,
					0.5,
					0.0,
					1.0,
					0.0029588245219072285,
					9.353787129822819e-14,
				),
			)
		})
	})
	Context("differentiation", func() {
		It("should correctly apply logistic differentiation to vector", func() {
			test.Equate(
				neuron.Differentiate(
					test.Vector(0.111293, 1, 0, 0.373892, 0.45, 0.1, 1.491e-12, 0.8, 0.0089144, 1.046e-7),
				),
				test.Vector(
					0.098906868151,
					0,
					0,
					0.23409677233600001,
					0.24750000000000003,
					0.09000000000000001,
					1.490999999997777e-12,
					0.15999999999999998,
					0.00883493347264,
					1.0459998905884e-7,
				),
			)
		})
	})
})
