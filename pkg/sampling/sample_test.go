package sampling_test

import (
	"github.com/google/go-cmp/cmp"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
)

var _ = Describe("sample", func() {
	Context("creation", func() {
		It("should construct a normal struct", func() {

			_ = sampling.NewSample(mat.NewVecDense(2, []float64{0.1, -0.2}), 1)
		})
		It("should construct a struct with zero activations", func() {

			_ = sampling.NewSample(mat.NewVecDense(6, nil), -4)
		})
		It("shouldn't construct `cause of nil activations", func() {
			defer ExpectPanic()
			_ = sampling.NewSample(nil, 6)
		})
	})
	Context("comparison", func() {
		It("should equal itself", func() {

			sample := sampling.NewSample(mat.NewVecDense(4, []float64{0.1, 0.2, -0.3, 0}), 4)
			Expect(cmp.Equal(sample, sample)).To(BeTrue())
		})
		It("shouldn't equal nil", func() {

			Expect(cmp.Equal(sampling.NewSample(mat.NewVecDense(2, []float64{0.2, -0.54}), 6), nil)).To(BeFalse())
		})
		It("shouldn't equal `cause of different activations", func() {

			Expect(
				cmp.Equal(
					sampling.NewSample(mat.NewVecDense(4, []float64{1, 2, -3, 7}), 1),
					sampling.NewSample(mat.NewVecDense(4, []float64{1, 0.2, -1, 0.61}), 1),
				),
			).To(BeFalse())
		})
		It("shouldn't equal `cause of different labels", func() {

			Expect(
				cmp.Equal(
					sampling.NewSample(mat.NewVecDense(2, []float64{0.167, 2}), 1),
					sampling.NewSample(mat.NewVecDense(2, []float64{0.167, 2}), 9),
				),
			).To(BeFalse())
		})
		It("shouldn't equal `cause of total difference", func() {

			Expect(
				cmp.Equal(
					sampling.NewSample(mat.NewVecDense(2, []float64{0.45, 2.08}), 1),
					sampling.NewSample(mat.NewVecDense(5, []float64{-1, -2, 0.3, 0.7, -1.06}), 12),
				),
			).To(BeFalse())
		})
		It("should equal with shared vector", func() {

			activations := mat.NewVecDense(3, []float64{1, 2, -3})
			Expect(cmp.Equal(sampling.NewSample(activations, 1), sampling.NewSample(activations, 1))).To(BeTrue())
		})
		It("should equal with equal-by-value vectors", func() {

			Expect(
				cmp.Equal(
					sampling.NewSample(mat.NewVecDense(3, []float64{1, 2, -3}), 1),
					sampling.NewSample(mat.NewVecDense(3, []float64{1, 2, -3}), 1),
				),
			).To(BeTrue())
		})
	})
})
