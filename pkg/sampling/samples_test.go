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
		//It("shouldn't construct `cause of elements with different activation lengths", func() {})
		//It("shouldn't construct `cause of elements with nil activations", func() {})
		//It("shouldn't construct `cause of trash elements", func() {})
		//It("should construct a collection with multiple robust samples", func() {})
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
