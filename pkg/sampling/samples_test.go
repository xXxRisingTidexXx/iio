package sampling_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/onsi/ginkgo"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
	"iio/pkg/test"
	"math/rand"
)

var _ = ginkgo.Describe("samples", func() {
	ginkgo.Context("creation", func() {
		test.With("should construct from nil slice", func() {
			test.Equate(sampling.NewSamples().Length(), 0)
		})
		test.With("should construct from zero-length slice", func() {
			test.Equate(sampling.NewSamples(make([]*sampling.Sample, 0)...).Length(), 0)
		})
		test.Spare("shouldn't construct `cause of single nil variadic elements", func() {
			_ = sampling.NewSamples(nil)
		})
		test.Spare("shouldn't construct `cause of nil variadic elements", func() {
			_ = sampling.NewSamples(nil, nil)
		})
		test.Spare("shouldn't construct `cause of slice test.With implicit nils", func() {
			_ = sampling.NewSamples(make([]*sampling.Sample, 10)...)
		})
		test.Spare("shouldn't construct `cause of slice test.With explicit nils", func() {
			_ = sampling.NewSamples([]*sampling.Sample{nil, nil, nil}...)
		})
		test.With("should construct from a single non-nil element", func() {
			sample := sampling.NewSample(mat.NewVecDense(4, []float64{3.2, -5.7, 0.1, -4}), 5)
			samples := sampling.NewSamples(sample)
			test.Equate(samples.Length(), 1)
			test.Equate(samples.Get(0), sample)
		})
		test.Spare("shouldn't construct `cause of elements test.With different activation lengths", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(4, []float64{0.34, 0.568, 0.981, 0.002}), 3),
				sampling.NewSample(mat.NewVecDense(4, []float64{0.5, 0.6667, 0.758, 0.03}), 9),
				sampling.NewSample(mat.NewVecDense(6, []float64{0.403, 0.8, 0.1, 0.65, 0, 0.7}), 7),
			)
		})
		test.Spare("shouldn't construct `cause of trash elements", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0.2, 0.3}), 1),
				nil,
				sampling.NewSample(mat.NewVecDense(5, []float64{1, 2, 3, 4, 5}), 6),
			)
		})
		test.With("should construct a collection test.With zero samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
				sampling.NewSample(mat.NewVecDense(4, nil), 0),
			)
			test.Equate(samples.Length(), 4)
			test.Comply(samples.Get(0), samples.Get(1))
			test.Comply(samples.Get(1), samples.Get(2))
			test.Comply(samples.Get(2), samples.Get(3))
		})
		test.With("should construct a collection test.With multiple robust samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(5, []float64{0.1, 0.2, 0.3, 0.3, 0.1}), 4),
				sampling.NewSample(mat.NewVecDense(5, []float64{0, 1, 1, 0.2, 0}), 5),
				sampling.NewSample(mat.NewVecDense(5, []float64{0.102, 0.4628, 0.21, 0.111, 0.97}), 2),
				sampling.NewSample(mat.NewVecDense(5, nil), 0),
				sampling.NewSample(mat.NewVecDense(5, []float64{0, 1, 1, 0.2, 0}), 5),
			)
			test.Equate(samples.Length(), 5)
			test.Comply(samples.Get(1), samples.Get(4))
		})
	})
	ginkgo.Context("comparison", func() {
		test.With("should equate the same-reference samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(5, nil), 0),
				sampling.NewSample(mat.NewVecDense(5, []float64{0, 1, 1, 0, 1}), 2),
			)
			test.Comply(samples, samples)
		})
		test.With("shouldn't equate nil and non-nil samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(5, nil), 0),
				sampling.NewSample(mat.NewVecDense(5, []float64{0.003, 0.98, 1, 0.6, 0.1}), 2),
			)
			test.Discern(samples, nil)
			test.Discern(nil, samples)
		})
		test.With("should equate samples from nil & non-nil slices", func() {
			test.Comply(sampling.NewSamples(), sampling.NewSamples(make([]*sampling.Sample, 0)...))
		})
		test.With("should equate non-empty variadic and slice-like samples", func() {
			test.Comply(
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(5, []float64{0.2, 1, 0.7, 1, 1}), 4),
				),
				sampling.NewSamples(
					[]*sampling.Sample{
						sampling.NewSample(mat.NewVecDense(5, []float64{0.2, 1, 0.7, 1, 1}), 4),
					}...,
				),
			)
		})
		test.With("shouldn't equate samples of different lengths", func() {
			test.Discern(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(4, []float64{0.29, 1, 0, 1}), 2)),
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0.86, 0, 1, 0.73}), 3),
					sampling.NewSample(mat.NewVecDense(4, []float64{1, 0.617, 0, 0.016}), 8),
				),
			)
		})
		test.With("shouldn't equate samples of different content", func() {
			test.Discern(
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 1, 0, 0.11}), 2),
					sampling.NewSample(mat.NewVecDense(4, nil), 0),
				),
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0.81, 0.9, 0.621, 0.3}), 3),
					sampling.NewSample(mat.NewVecDense(4, []float64{1, 0, 0, 0}), 8),
				),
			)
		})
		test.With("shouldn't equate `cause of different element order", func() {
			test.Discern(
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0.89, 0, 0.4}), 7),
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 1, 0, 0}), 1),
				),
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 1, 0, 0}), 1),
					sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0.89, 0, 0.4}), 7),
				),
			)
		})
		test.With("should equate the same-content samples", func() {
			test.Comply(
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 0, 0}), 6),
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 1, 1, 1}), 9),
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 0.45, 0.87, 0.1}), 8),
				),
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 0, 0}), 6),
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 1, 1, 1}), 9),
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 0.45, 0.87, 0.1}), 8),
				),
			)
		})
	})
	ginkgo.Context("slicing", func() {
		test.With("to (-inf; 0] on empty samples", func() {
			test.Comply(sampling.NewSamples().To(-3), sampling.NewSamples())
		})
		test.With("to 1 on empty samples", func() {
			test.Comply(sampling.NewSamples().To(1), sampling.NewSamples())
		})
		test.With("to [2; +inf) on empty samples", func() {
			test.Comply(sampling.NewSamples().To(3), sampling.NewSamples())
		})
		test.With("from (-inf; -1] on empty samples", func() {
			test.Comply(sampling.NewSamples().From(-1), sampling.NewSamples())
		})
		test.With("from 0 on empty samples", func() {
			test.Comply(sampling.NewSamples().From(0), sampling.NewSamples())
		})
		test.With("from [1; +inf) on empty samples", func() {
			test.Comply(sampling.NewSamples().From(2), sampling.NewSamples())
		})
		test.With("to (-inf; 0] on single-element samples", func() {
			test.Comply(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 1, 0}), 5)).To(0),
				sampling.NewSamples(),
			)
		})
		test.With("to 1 on single-element samples", func() {
			test.Comply(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 1, 0}), 5)).To(1),
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 1, 0}), 5)),
			)
		})
		test.With("to [2; +inf) on single-element samples", func() {
			test.Comply(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 1, 0}), 5)).To(3),
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(4, []float64{0.9, 0, 1, 0}), 5)),
			)
		})
		test.With("from (-inf; -1] on single-element samples", func() {
			test.Comply(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 1}), 1)).From(-4),
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 1}), 1)),
			)
		})
		test.With("from 0 on single-element samples", func() {
			test.Comply(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 1}), 1)).From(0),
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 1}), 1)),
			)
		})
		test.With("from [1; +inf) on single-element samples", func() {
			test.Comply(
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 1}), 1)).From(1),
				sampling.NewSamples(),
			)
		})
		test.With("to (-inf; 0] on multi-element samples", func() {
			test.Comply(
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
					sampling.NewSample(mat.NewVecDense(3, []float64{0.61, 0, 0.5}), 1),
					sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
				).To(-10),
				sampling.NewSamples(),
			)
		})
		test.With("to 1 on multi-element samples", func() {
			test.Comply(
				sampling.NewSamples(
					sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
					sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
					sampling.NewSample(mat.NewVecDense(3, []float64{0.61, 0, 0.5}), 1),
				).To(1),
				sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2)),
			)
		})
		test.With("to [2; len - 1] on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.61, 0, 0.5}), 1),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0.579, 0.1}), 4),
					).To(2),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.61, 0, 0.5}), 1),
					),
				),
			).To(BeTrue())
		})
		test.With("to len on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0.579, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.61, 0, 0.5}), 1),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
					).To(4),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0.579, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.61, 0, 0.5}), 1),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
					),
				),
			).To(BeTrue())
		})
		test.With("to [len + 1; +inf) on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
					).To(6),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
					),
				),
			).To(BeTrue())
		})
		test.With("from (-inf; -1] on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
					).From(-60),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.4, 0.01}), 3),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
					),
				),
			).To(BeTrue())
		})
		test.With("from 0 on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{1, 1, 0.9}), 6),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
					).From(0),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{1, 1, 0.9}), 6),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
					),
				),
			).To(BeTrue())
		})
		test.With("from [1; len - 2] on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.2, 0.25, 0.8}), 7),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
					).From(2),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
					),
				),
			).To(BeTrue())
		})
		test.With("from len - 1 on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.2, 0.25, 0.8}), 7),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
					).From(3),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
					),
				),
			).To(BeTrue())
		})
		test.With("from [len; +inf) on multi-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.1, 0, 0.1}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.2, 0.25, 0.8}), 7),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 2),
					).From(3),
					sampling.NewSamples(),
				),
			).To(BeTrue())
		})
	})
	ginkgo.Context("indexing", func() {
		test.Spare("should fail in empty collection `cause of negative index", func() {
			_ = sampling.NewSamples().Get(-2)
		})
		test.Spare("should fail in empty collection `cause of zero index", func() {
			_ = sampling.NewSamples().Get(0)
		})
		test.Spare("should fail in empty collection `cause of positive index", func() {
			_ = sampling.NewSamples().Get(20)
		})
		test.Spare("should fail in single-element collection `cause of negative index", func() {
			_ = sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.5, 0.1}), 2)).Get(-30)
		})
		test.With("should yield in multi-element collection the single element", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.5, 0.1}), 2)).Get(0),
					sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.5, 0.1}), 2),
				),
			).To(BeTrue())
		})
		test.Spare("should fail in single-element collection `cause of positive index", func() {
			_ = sampling.NewSamples(sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 0.9}), 4)).Get(6)
		})
		test.Spare("should fail in multi-element collection `cause of negative index", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(4, []float64{0, 0.0002, 0.9, 0.107}), 7),
				sampling.NewSample(mat.NewVecDense(4, []float64{0.8, 0.2, 0.12, 0.7}), 6),
			).Get(-3)
		})
		test.With("should yield in multi-element collection the first element", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(4, []float64{0.782, 0.2, 0.1, 0}), 2),
						sampling.NewSample(mat.NewVecDense(4, []float64{0.6154, 0.8788, 0, 0}), 5),
						sampling.NewSample(mat.NewVecDense(4, nil), 0),
					).Get(0),
					sampling.NewSample(mat.NewVecDense(4, []float64{0.782, 0.2, 0.1, 0}), 2),
				),
			).To(BeTrue())
		})
		test.With("should yield in multi-element collection a middle element", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(4, []float64{0.52, 0.2, 0.1, 0}), 2),
						sampling.NewSample(mat.NewVecDense(4, nil), 0),
						sampling.NewSample(mat.NewVecDense(4, []float64{0, 0, 0, 0.87}), 3),
					).Get(1),
					sampling.NewSample(mat.NewVecDense(4, []float64{0, 0, 0, 0}), 0),
				),
			).To(BeTrue())
		})
		test.With("should yield in multi-element collection the last element", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(4, nil), 0),
						sampling.NewSample(mat.NewVecDense(4, []float64{0.2, 0, 0.1, 0}), 2),
						sampling.NewSample(mat.NewVecDense(4, []float64{0.52, 0.2, 0.1, 0}), 2),
						sampling.NewSample(mat.NewVecDense(4, []float64{1, 0, 1, 0.754}), 3),
					).Get(3),
					sampling.NewSample(mat.NewVecDense(4, []float64{1, 0, 1, 0.754}), 3),
				),
			).To(BeTrue())
		})
		test.Spare("should fail in multi-element collection `cause of too large index", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(4, []float64{0.8, 0.2, 0.12, 0.7}), 6),
				sampling.NewSample(mat.NewVecDense(4, []float64{0, 0.0002, 0.9, 0.107}), 7),
			).Get(109)
		})
	})
	ginkgo.Context("shuffling", func() {
		test.With("should yield the same empty samples", func() {
			Expect(cmp.Equal(sampling.NewSamples().Shuffle(), sampling.NewSamples())).To(BeTrue())
		})
		test.With("should yield the same single-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.8, 0.2, 0.1282}), 9),
					).Shuffle(),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.8, 0.2, 0.1282}), 9),
					),
				),
			).To(BeTrue())
		})
		test.With("should yield the same equal-element samples", func() {
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.5, 0.2, 0.3}), 8),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.5, 0.2, 0.3}), 8),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.5, 0.2, 0.3}), 8),
					).Shuffle(),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.5, 0.2, 0.3}), 8),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.5, 0.2, 0.3}), 8),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.5, 0.2, 0.3}), 8),
					),
				),
			).To(BeTrue())
		})
		test.With("should yield different multi-element samples", func() {
			rand.Seed(42)
			Expect(
				cmp.Equal(
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.8112, 0.2301, 0.6748}), 8),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.501293, 0.212893, 0.30231}), 9),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.123891, 0.93812, 0.30128}), 4),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.342340123, 1, 0.23923}), 7),
						sampling.NewSample(mat.NewVecDense(3, []float64{1, 0.2327485, 0.11192}), 1),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.6345, 0.00203}), 5),
					).Shuffle(),
					sampling.NewSamples(
						sampling.NewSample(mat.NewVecDense(3, []float64{0.501293, 0.212893, 0.30231}), 9),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.342340123, 1, 0.23923}), 7),
						sampling.NewSample(mat.NewVecDense(3, []float64{1, 0.2327485, 0.11192}), 1),
						sampling.NewSample(mat.NewVecDense(3, []float64{0, 0.6345, 0.00203}), 5),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.8112, 0.2301, 0.6748}), 8),
						sampling.NewSample(mat.NewVecDense(3, []float64{0.123891, 0.93812, 0.30128}), 4),
					),
				),
			).To(BeTrue())
		})
	})
	ginkgo.Context("batching", func() {
		test.Spare("shouldn't straightly batch over empty samples", func() {
			_ = sampling.NewSamples().Batch(1)
		})
		test.Spare("shouldn't batch test.With negative size", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.7593, 0.01028, 0.9117}), 0),
			).Batch(-4)
		})
		test.Spare("shouldn't batch test.With zero size", func() {
			_ = sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.89012, 0.999, 1}), 1),
			).Batch(0)
		})
		test.With("should batch single-element samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.3902, 0.00023, 0.0045}), 1),
			)
			batch := samples.Batch(1)
			Expect(batch.Length()).To(And(Equal(1), Equal(samples.Length())))
			Expect(cmp.Equal(samples.Get(0), batch.Get(0))).To(BeTrue())
		})
		test.With("should batch test.With full size over multi-element samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.3902, 0.00023, 0.0045}), 1),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.675483, 0.123, 0.75849}), 2),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.878685, 0.00123, 1}), 3),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.444563, 0.1123, 0.90134}), 4),
			)
			batch1 := samples.Batch(2)
			batch2 := samples.Batch(2)
			Expect(2).To(And(Equal(batch1.Length()), Equal(batch2.Length())))
			Expect(cmp.Equal(batch1.Get(0), samples.Get(0))).To(BeTrue())
			Expect(cmp.Equal(batch1.Get(1), samples.Get(1))).To(BeTrue())
			Expect(cmp.Equal(batch2.Get(0), samples.Get(2))).To(BeTrue())
			Expect(cmp.Equal(batch2.Get(1), samples.Get(3))).To(BeTrue())
		})
		test.With("should batch test.With partial size over multi-element samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.675483, 0.123, 0.75849}), 2),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.3902, 0.00023, 0.0045}), 1),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.444563, 0.1123, 0.90134}), 4),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.878685, 0.00123, 1}), 3),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.239987, 0.00001, 0.9762}), 3),
			)
			batch1 := samples.Batch(3)
			batch2 := samples.Batch(3)
			Expect(batch1.Length()).To(Equal(3))
			Expect(batch2.Length()).To(Equal(2))
			Expect(cmp.Equal(batch1.Get(0), samples.Get(0))).To(BeTrue())
			Expect(cmp.Equal(batch1.Get(1), samples.Get(1))).To(BeTrue())
			Expect(cmp.Equal(batch1.Get(2), samples.Get(2))).To(BeTrue())
			Expect(cmp.Equal(batch2.Get(0), samples.Get(3))).To(BeTrue())
			Expect(cmp.Equal(batch2.Get(1), samples.Get(4))).To(BeTrue())
		})
		test.Spare("shouldn't batch more then once in a row", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.878685, 0.00123, 1}), 3),
				sampling.NewSample(mat.NewVecDense(3, []float64{0.239987, 0.00001, 0.9762}), 3),
			)
			batch := samples.Batch(2)
			Expect(batch.Length()).To(Equal(2))
			Expect(cmp.Equal(batch.Get(0), samples.Get(0))).To(BeTrue())
			Expect(cmp.Equal(batch.Get(1), samples.Get(1))).To(BeTrue())
			_ = samples.Batch(2)
		})
		test.With("should distinct partially-iterated and fully-iterated samples", func() {
			samples1 := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.239987, 1, 1}), 7),
				sampling.NewSample(mat.NewVecDense(3, nil), 5),
				sampling.NewSample(mat.NewVecDense(3, nil), 0),
			)
			samples2 := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0.239987, 1, 1}), 7),
				sampling.NewSample(mat.NewVecDense(3, nil), 5),
				sampling.NewSample(mat.NewVecDense(3, nil), 0),
			)
			Expect(cmp.Equal(samples1, samples2)).To(BeTrue())
			for samples1.Next() {
				_ = samples1.Batch(2)
				Expect(cmp.Equal(samples1, samples2)).To(BeFalse())
			}
			Expect(cmp.Equal(samples1, samples2)).To(BeTrue())
		})
		test.With("should normally iterate test.With multiple batch size", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 0}), 0),
				sampling.NewSample(mat.NewVecDense(3, []float64{0, 0, 1}), 1),
				sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 0}), 2),
				sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 3),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 0, 0}), 4),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 0, 1}), 5),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 1, 0}), 6),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 1, 1}), 7),
			)
			for i := 0; samples.Next(); i++ {
				batch := samples.Batch(3)
				for j := 0; j < batch.Length(); j++ {
					Expect(cmp.Equal(batch.Get(j), samples.Get(i*3+j))).To(BeTrue())
				}
			}
		})
		test.With("should normally reiterate over the same samples", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(2, []float64{0, 0}), 0),
				sampling.NewSample(mat.NewVecDense(2, []float64{0, 1}), 1),
				sampling.NewSample(mat.NewVecDense(2, []float64{1, 0}), 2),
				sampling.NewSample(mat.NewVecDense(2, []float64{1, 1}), 3),
			)
			for samples.Next() {
				Expect(samples.Batch(2).Length()).To(Equal(2))
			}
			Expect(samples.Next()).To(BeTrue())
			for samples.Next() {
				Expect(samples.Batch(2).Length()).To(Equal(2))
			}
			Expect(samples.Next()).To(BeTrue())
		})
	})
	ginkgo.Context("scenarios", func() {
		test.With("should split the samples and get a few elements", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(2, []float64{1, 1}), 3),
				sampling.NewSample(mat.NewVecDense(2, []float64{1, 0}), 2),
				sampling.NewSample(mat.NewVecDense(2, []float64{0, 1}), 1),
				sampling.NewSample(mat.NewVecDense(2, []float64{0, 0}), 0),
			)
			left, right := samples.To(2), samples.From(2)
			Expect(cmp.Equal(left.Get(0), samples.Get(0))).To(BeTrue())
			Expect(cmp.Equal(left.Get(1), samples.Get(1))).To(BeTrue())
			Expect(cmp.Equal(right.Get(0), samples.Get(2))).To(BeTrue())
			Expect(cmp.Equal(right.Get(1), samples.Get(3))).To(BeTrue())
		})
		test.With("should adequately leverage slicing", func() {
			samples := sampling.NewSamples(
				sampling.NewSample(mat.NewVecDense(3, []float64{0, 1, 1}), 3),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 1, 2}), 1),
				sampling.NewSample(mat.NewVecDense(3, []float64{1, 2, 3}), 2),
				sampling.NewSample(mat.NewVecDense(3, []float64{2, 3, 5}), 0),
				sampling.NewSample(mat.NewVecDense(3, []float64{3, 5, 8}), 4),
				sampling.NewSample(mat.NewVecDense(3, []float64{5, 8, 13}), 6),
				sampling.NewSample(mat.NewVecDense(3, []float64{8, 13, 21}), 7),
				sampling.NewSample(mat.NewVecDense(3, []float64{13, 21, 34}), 8),
				sampling.NewSample(mat.NewVecDense(3, []float64{21, 34, 55}), 9),
				sampling.NewSample(mat.NewVecDense(3, []float64{34, 55, 89}), 5),
			)
			Expect(
				cmp.Equal(
					samples.From(2).To(7).To(6).To(2).From(1).Get(0),
					sampling.NewSample(mat.NewVecDense(3, []float64{2, 3, 5}), 0),
				),
			).To(BeTrue())
			Expect(
				cmp.Equal(
					samples.To(20).From(1).To(7).To(5).To(3).From(-3).To(2).From(3),
					sampling.NewSamples(),
				),
			).To(BeTrue())
		})
	})
})
