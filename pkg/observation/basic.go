package observation

import (
	"gonum.org/v1/gonum/mat"
	"sync"
)

func NewBasicObserver(options *Options) Observer {
	if options == nil {
		panic("observation: basic observer got nil options")
	}
	bucketNumber := options.SetLength / options.PortionSize
	lastPortionSize := options.SetLength % options.PortionSize
	if lastPortionSize != 0 {
		bucketNumber++
	} else {
		lastPortionSize = options.PortionSize
	}
	length := options.EpochNumber * bucketNumber
	return &basicObserver{
		buckets:         mat.NewVecDense(length, nil),
		length:          length,
		bucketNumber:    bucketNumber,
		portionSize:     options.PortionSize,
		lastPortionSize: lastPortionSize,
	}
}

type basicObserver struct {
	sync.Mutex
	buckets          *mat.VecDense
	length           int
	bucketNumber     int
	portionSize      int
	lastPortionSize  int
	bucketIndex      int
	observationIndex int
}

func (observer *basicObserver) Observe(cost float64) {
	observer.Lock()
	bucketIndex := observer.bucketIndex
	if bucketIndex >= observer.length {
		panic("observe: basic observer observation ended")
	}
	observer.observationIndex++
	value := observer.buckets.AtVec(bucketIndex) + cost
	portionSize := observer.portionSize
	if (bucketIndex+1)%observer.bucketNumber == 0 {
		portionSize = observer.lastPortionSize
	}
	if observer.observationIndex == portionSize {
		value /= float64(portionSize)
		observer.observationIndex = 0
		observer.bucketIndex++
	}
	observer.buckets.SetVec(bucketIndex, value)
	observer.Unlock()
}

func (observer *basicObserver) Expound() mat.Matrix {
	observations := mat.NewDense(observer.length, 2, nil)
	for i := 0; i < observer.length; i++ {
		observations.Set(i, 0, float64(i+1)/float64(observer.bucketNumber))
		observations.Set(i, 1, observer.buckets.AtVec(i))
	}
	return observations
}
