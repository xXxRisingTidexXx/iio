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
			estimate.NewReport(
				[]*estimate.Record{
					estimate.NewRecord(1, 0.5, 1, 0.6666666666666666),
					estimate.NewRecord(1, 0, 0, 0),
					estimate.NewRecord(3, 1, 0.6666666666666666, 0.8),
				},
				estimate.NewRecord(5, 0.5, 0.5555555555555555, 0.48888888888888893),
				0.6,
			),
		)
	})
})
