package sampling_test

import (
	"github.com/google/go-cmp/cmp"
	"iio/pkg/sampling"
)

var _ = Describe("samples", func() {
	Context("creation", func() {
		It("should construct from nil slice", func() {
			defer ExpectNoPanic()
			Expect(sampling.NewSamples().Length()).To(Equal(0))
		})
		It("should construct from zero-length slice", func() {
			defer ExpectNoPanic()
			Expect(sampling.NewSamples(make([]*sampling.Sample, 0)...).Length()).To(Equal(0))
		})
		It("shouldn't construct `cause of single nil variadic elements", func() {
			defer ExpectPanic()
			_ = sampling.NewSamples(nil)
		})
		It("shouldn't construct `cause of nil variadic elements", func() {
			defer ExpectPanic()
			_ = sampling.NewSamples(nil, nil)
		})
		It("shouldn't construct `cause of slice with implicit nils", func() {
			defer ExpectPanic()
			_ = sampling.NewSamples(make([]*sampling.Sample, 10)...)
		})
		It("shouldn't construct `cause of slice with explicit nils", func() {
			defer ExpectPanic()
			_ = sampling.NewSamples([]*sampling.Sample{nil, nil, nil}...)
		})
		//It("should construct from a single non-nil element", func() {
		//	defer expectNoPanic()
		//
		//})
		//It("shouldn't construct `cause of elements with different activation lengths", func() {})
		//It("shouldn't construct `cause of elements with nil activations", func() {})
		//It("shouldn't construct `cause of trash elements", func() {})
		//It("should construct a collection with multiple robust samples", func() {})
	})
	Context("comparison", func() {
		//It("should equate the same samples", func() {})
		//It("shouldn't equate nil and non-nil samples", func() {})
		It("should equate samples from nil & non-nil slices", func() {
			defer ExpectNoPanic()
			Expect(cmp.Equal(sampling.NewSamples(), sampling.NewSamples(make([]*sampling.Sample, 0)...))).To(BeTrue())
		})
		//It("should equate non-empty variadic and slice-like samples", func() {})
		//It("shouldn't equate samples of different lengths", func() {})
		//It("shouldn't equate iterating and non-iterating samples", func() {})
		//It("should equate already-iterated and non-iterated samples", func() {})
	})
	Context("slicing", func() {})
	Context("indexing", func() {})
	Context("shuffling", func() {})
	Context("batching", func() {})
	Context("scenarios", func() {})
})
