package observation

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"sync"
)

func NewBasicObserver(epochNumber, setLength, portionSize int) *BasicObserver {
	if epochNumber < 1 {
		panic(fmt.Sprintf("observer: basic observer got invalid epoch number, %d", epochNumber))
	}
	if setLength < 1 {
		panic(fmt.Sprintf("observer: basic observer got invalid set length, %d", setLength))
	}
	if portionSize < 1 {
		panic(fmt.Sprintf("observer: basic observer got invalid portion size, %d", portionSize))
	}
	epochBucketNumber, lastPortionSize := setLength/portionSize, setLength%portionSize
	if lastPortionSize != 0 {
		epochBucketNumber++
	} else {
		lastPortionSize = portionSize
	}
	length := epochNumber*epochBucketNumber
	return &BasicObserver{
		buckets:           mat.NewVecDense(length, nil),
		length:            length,
		epochBucketNumber: epochBucketNumber,
		portionSize:       portionSize,
		lastPortionSize:   lastPortionSize,
	}
}

type BasicObserver struct {
	sync.Mutex
	buckets           *mat.VecDense
	length            int
	epochBucketNumber int
	portionSize       int
	lastPortionSize   int
	bucketIndex       int
	observationIndex  int
}

func (observer *BasicObserver) Observe(cost float64) {
	observer.Lock()
	bucketIndex := observer.bucketIndex
	if bucketIndex >= observer.length {
		panic("observe: basic observer observation ended")
	}
	observer.observationIndex++
	value := observer.buckets.AtVec(bucketIndex) + cost
	portionSize := observer.portionSize
	if (bucketIndex+1)%observer.epochBucketNumber == 0 {
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

func (observer *BasicObserver) Expound() mat.Matrix {
	return observer.buckets
}
