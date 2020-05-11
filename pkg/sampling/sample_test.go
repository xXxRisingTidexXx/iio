package sampling_test

import (
	"github.com/onsi/ginkgo"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
	"iio/pkg/test"
)

var _ = ginkgo.Describe("sample", func() {
	ginkgo.Context("creation", func() {
		test.With("should construct a normal struct", func() {
			_ = sampling.NewSample(mat.NewVecDense(2, []float64{0.1, -0.2}), 1)
		})
		test.With("should construct a struct test.With zero activations", func() {
			_ = sampling.NewSample(mat.NewVecDense(6, nil), -4)
		})
		test.Spare("shouldn't construct `cause of nil activations", func() {
			_ = sampling.NewSample(nil, 6)
		})
	})
	ginkgo.Context("comparison", func() {
		test.With("should equal itself", func() {
			sample := sampling.NewSample(mat.NewVecDense(4, []float64{0.1, 0.2, -0.3, 0}), 4)
			test.Equate(sample, sample)
		})
		test.With("shouldn't equal nil", func() {
			test.Discern(sampling.NewSample(mat.NewVecDense(2, []float64{0.2, -0.54}), 6), nil)
		})
		test.With("shouldn't equal `cause of different activations", func() {
			test.Discern(
				sampling.NewSample(mat.NewVecDense(4, []float64{1, 2, -3, 7}), 1),
				sampling.NewSample(mat.NewVecDense(4, []float64{1, 0.2, -1, 0.61}), 1),
			)
		})
		test.With("shouldn't equal `cause of different labels", func() {
			test.Discern(
				sampling.NewSample(mat.NewVecDense(2, []float64{0.167, 2}), 1),
				sampling.NewSample(mat.NewVecDense(2, []float64{0.167, 2}), 9),
			)
		})
		test.With("shouldn't equal `cause of total difference", func() {
			test.Discern(
				sampling.NewSample(mat.NewVecDense(2, []float64{0.45, 2.08}), 1),
				sampling.NewSample(mat.NewVecDense(5, []float64{-1, -2, 0.3, 0.7, -1.06}), 12),
			)
		})
		test.With("should equal with shared vector", func() {
			activations := mat.NewVecDense(3, []float64{1, 2, -3})
			test.Equate(sampling.NewSample(activations, 1), sampling.NewSample(activations, 1))
		})
		test.With("should equal with equal-by-value vectors", func() {
			test.Equate(
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 2, -3}), 1),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 2, -3}), 1),
			)
		})
	})
})
