package sampling_test

import (
	"github.com/google/go-cmp/cmp"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
)

var _ = Describe("samples", func() {
	Context("creation", func() {
		With("should construct from nil slice", func() {
			Expect(sampling.NewSamples().Length()).To(Equal(0))
		})
		With("should construct from zero-length slice", func() {
			Expect(sampling.NewSamples(make([]*sampling.Sample, 0)...).Length()).To(Equal(0))
		})
		Spare("shouldn't construct `cause of single nil variadic elements", func() {
			_ = sampling.NewSamples(nil)
		})
		Spare("shouldn't construct `cause of nil variadic elements", func() {
			_ = sampling.NewSamples(nil, nil)
		})
		Spare("shouldn't construct `cause of slice with implicit nils", func() {
			_ = sampling.NewSamples(make([]*sampling.Sample, 10)...)
		})
		Spare("shouldn't construct `cause of slice with explicit nils", func() {
			_ = sampling.NewSamples([]*sampling.Sample{nil, nil, nil}...)
		})
		With("should construct from a single non-nil element", func() {
			sample := sampling.NewSample(mat.NewVecDense(4, []float64{3.2, -5.7, 0.1, -4}), 5)
			samples := sampling.NewSamples(sample)
			Expect(samples.Length()).To(Equal(1))
			Expect(samples.Get(0)).To(Equal(sample))
		})
		Spare("shouldn't construct `cause of elements with different activation lengths", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(4, []float64{0.34, 0.568, 0.981, 0.002}), 3),
				sampling.NewSample(mat.NewVecDense(4, []float64{0.5, 0.6667, 0.758, 0.03}), 9),
				sampling.NewSample(mat.NewVecDense(6, []float64{0.403, 0.8, 0.1, 0.65, 0, 0.7}), 7),
			)
		})
		Spare("shouldn't construct `cause of trash elements", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0.2, 0.3}), 1),
				nil,
				sampling.NewSample(mat.NewVecDense(5, []float64{1, 2, 3, 4, 5}), 6),
			)
		})
		With("should construct a collection with zero samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
			)
			Expect(samples.Length()).To(Equal(4))
			Expect(cmp.Equal(samples.Get(0), samples.Get(1))).To(BeTrue())
			Expect(cmp.Equal(samples.Get(1), samples.Get(2))).To(BeTrue())
			Expect(cmp.Equal(samples.Get(2), samples.Get(3))).To(BeTrue())
		})
		//With("should construct a collection with multiple robust samples", func() {})
	})
	Context("comparison", func() {
		//It("should equate the same samples", func() {})
		//It("shouldn't equate nil and non-nil samples", func() {})
		With("should equate samples from nil & non-nil slices", func() {
			Expect(cmp.Equal(sampling.NewSamples(), sampling.NewSamples(make([]*sampling.Sample, 0)...))).
				To(BeTrue())
		})
		//It("should equate non-empty variadic and slice-like samples", func() {})
		//It("shouldn't equate samples of different lengths", func() {})
	})
	Context("slicing", func() {})
	Context("indexing", func() {})
	Context("shuffling", func() {})
	Context("batching", func() {
		//It("shouldn't equate iterating and non-iterating samples", func() {})
		//It("should equate already-iterated and non-iterated samples", func() {})
	})
	Context("scenarios", func() {})
})
