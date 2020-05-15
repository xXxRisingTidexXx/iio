package estimate

import (
	"fmt"
	"sync"
)

func NewBasicEstimator(classNumber int) *BasicEstimator {
	if classNumber <= 1 {
		panic(fmt.Sprintf("reports: got invalid class number, %d", classNumber))
	}
	confusionMatrix := make([][]int, classNumber)
	for i := range confusionMatrix {
		confusionMatrix[i] = make([]int, classNumber)
	}
	return &BasicEstimator{classNumber: classNumber, confusionMatrix: confusionMatrix}
}

type BasicEstimator struct {
	sync.RWMutex
	classNumber     int
	confusionMatrix [][]int
}

func (reporter *BasicEstimator) Track(actual, ideal int) {
	if actual < 0 || actual >= reporter.classNumber {
		panic(fmt.Sprintf("reports: invalid actual label, %d", actual))
	}
	if ideal < 0 || ideal >= reporter.classNumber {
		panic(fmt.Sprintf("reports: invalid ideal label, %d", ideal))
	}
	reporter.Lock()
	reporter.confusionMatrix[actual][ideal]++
	reporter.Unlock()
}

func (reporter *BasicEstimator) Estimate() *Report {
	panic("implement me")
}
