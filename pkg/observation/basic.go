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
	bucketNumber, lastPortionSize := setLength/portionSize, setLength%portionSize
	if lastPortionSize != 0 {
		bucketNumber++
	} else {
		lastPortionSize = portionSize
	}
	length := epochNumber * bucketNumber
	return &BasicObserver{
		buckets:         mat.NewVecDense(length, nil),
		length:          length,
		bucketNumber:    bucketNumber,
		portionSize:     portionSize,
		lastPortionSize: lastPortionSize,
	}
}

type BasicObserver struct {
	sync.Mutex
	buckets          *mat.VecDense
	length           int
	bucketNumber     int
	portionSize      int
	lastPortionSize  int
	bucketIndex      int
	observationIndex int
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

func (observer *BasicObserver) Expound() mat.Matrix {
	observations := mat.NewDense(observer.length+1, 2, nil)
	observations.Set(0, 0, 0)
	observations.Set(0, 1, 0)
	for i := 1; i <= observer.length; i++ {
		observations.Set(i, 0, float64(i)/float64(observer.bucketNumber))
		observations.Set(i, 1, observer.buckets.AtVec(i-1))
	}
	return observations
}
