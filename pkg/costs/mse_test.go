package costs_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/costs"
	"iio/pkg/test"
)

var _ = Describe("mse", func() {
	costFunction := costs.NewMSECostFunction()
	Context("evaluation", func() {
		test.With("should correctly calculate the final class among two equal ones", func() {
			test.Comply(
				costFunction.Evaluate(test.Vector(0, 0.003, 0.027, 0.21, 0.308, 0.308, 0.0004, 0.00037)),
				4,
			)
		})
	})
	Context("cost", func() {
		test.With("should correctly calculate the final output cost", func() {
			test.Comply(
				costFunction.Cost(test.Vector(0.001, 0.00032, 0.472, 0.1102, 0.63118), 4),
				0.37095733480000004,
			)
		})
	})
	Context("differentiation", func() {
		test.With("should correctly calculate the differential", func() {
			test.Equate(
				costFunction.Differentiate(test.Vector(0.0004, 0.10056, 0.00382, 0.302, 0.11134, 0.762), 3),
				test.Vector(0.0004, 0.10056, 0.00382, -0.698, 0.11134, 0.762),
			)
		})
	})
})
