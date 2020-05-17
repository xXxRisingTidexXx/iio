package observation_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/observation"
	"iio/pkg/test"
	"math/rand"
	"sync"
	"time"
)

var _ = Describe("basic", func() {
	observe := func(observer observation.Observer, cost float64, waitGroup *sync.WaitGroup) {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		observer.Observe(cost)
		waitGroup.Done()
	}
	test.With("should correctly collect all costs and produce a series", func() {
		rand.Seed(time.Now().UnixNano())
		observer := observation.NewBasicObserver(2, 10, 3)
		costs := []float64{
			0.94083628,
			0.82872992,
			1.42810471,
			1.33250802,
			0.79602682,
			1.48335108,
			0.64320754,
			0.92900801,
			0.88629881,
			1.24980291,
			0.92818381,
			0.86987584,
			0.92233395,
			1.68346876,
			-0.15947398,
			0.72776115,
			1.43894503,
			1.11551173,
			0.88204915,
			0.63063128,
		}
		waitGroup := &sync.WaitGroup{}
		for i := 0; i < 8; i++ {
			start, offset := i/4*10+i%4*3, 3
			if i%4 == 3 {
				offset = 1
			}
			batch := costs[start : start+offset]
			waitGroup.Add(len(batch))
			for _, cost := range batch {
				go observe(observer, cost, waitGroup)
			}
			waitGroup.Wait()
		}
		test.Equate(
			observer.Expound(),
			test.Matrix(
				9,
				2,
				0,
				0,
				0.25,
				1.0658903033333333,
				0.5,
				1.2039619733333333,
				0.75,
				0.8195047866666667,
				1,
				1.24980291,
				1.25,
				0.9067978666666666,
				1.5,
				0.75058531,
				1.75,
				1.14550197,
				2,
				0.63063128,
			),
		)
	})
})
