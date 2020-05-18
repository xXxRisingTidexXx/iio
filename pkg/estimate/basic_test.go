package estimate_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/estimate"
	"iio/pkg/test"
	"math/rand"
	"sync"
	"time"
)

var _ = Describe("basic", func() {
	track := func(estimator estimate.Estimator, actual, ideal int, waitGroup *sync.WaitGroup) {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		estimator.Track(actual, ideal)
		waitGroup.Done()
	}
	test.With("should correctly track & estimate all predictions", func() {
		rand.Seed(time.Now().UnixNano())
		estimator := estimate.NewBasicEstimator(estimate.NewOptions(3))
		actuals, ideals := []int{0, 0, 2, 2, 1}, []int{0, 1, 2, 2, 2}
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(5)
		for i := 0; i < 5; i++ {
			go track(estimator, actuals[i], ideals[i], waitGroup)
		}
		waitGroup.Wait()
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
