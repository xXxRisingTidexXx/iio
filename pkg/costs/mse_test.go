package costs_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"iio/pkg/costs"
	"iio/pkg/test"
)

var _ = Describe("mse", func() {
	costFunction := costs.NewMSECostFunction()
	Context("evaluation", func() {
		It("should correctly calculate the final class among two equal ones", func() {
			gomega.Expect(
				costFunction.Evaluate(test.Vector(0, 0.003, 0.027, 0.21, 0.308, 0.308, 0.0004, 0.00037)),
			).To(gomega.Equal(4))
		})
	})
	Context("differentiation", func() {
		It("should correctly calculate the differential", func() {
			test.Equate(
				costFunction.Differentiate(test.Vector(0.0004, 0.10056, 0.00382, 0.302, 0.11134, 0.762), 3),
				test.Vector(0.0004, 0.10056, 0.00382, -0.698, 0.11134, 0.762),
			)
		})
	})
})
