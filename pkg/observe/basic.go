package observe

import (
	"gonum.org/v1/gonum/mat"
	"sync"
)

func NewBasicObserver(epochNumber, setLength, batchSize int) *BasicObserver {
	return &BasicObserver{}
}

type BasicObserver struct {
	sync.Mutex
}

func (observer *BasicObserver) Observe(value float64) {
	panic("implement me")
}

func (observer *BasicObserver) Expound() mat.Matrix {
	panic("implement me")
}

