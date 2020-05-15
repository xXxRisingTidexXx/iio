package estimate_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/estimate"
	"iio/pkg/test"
)

var _ = Describe("basic", func() {
	test.With("should correctly track & estimate all predictions", func() {
		estimator := estimate.NewBasicEstimator(3)
		estimator.Track(0, 0)
		estimator.Track(0, 1)
		estimator.Track(2, 2)
		estimator.Track(2, 2)
		estimator.Track(1, 2)
		test.Comply(
			estimator.Estimate(),
			&estimate.Report{
				Classes: []*estimate.Record{
					{0.5, 1, 0.6666666666666666, 1},
					{0, 0, 0, 1},
					{1, 0.6666666666666666, 0.8, 3},
				},
				MacroAvg: &estimate.Record{
					Precision: 0.5,
					Recall:    0.5555555555555555,
					F1Score:   0.48888888888888893,
					Support:   5,
				},
				Accuracy: 0.6,
			},
		)
	})
})
