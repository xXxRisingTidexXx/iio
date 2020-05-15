package reports

import (
	"fmt"
	"sync"
)

func NewBasicReporter(classNumber int) *BasicReporter {
	if classNumber <= 1 {
		panic(fmt.Sprintf("reports: got invalid class number, %d", classNumber))
	}

	return &BasicReporter{}
}

type BasicReporter struct {
	sync.RWMutex
	classNumber int
	confusionMatrix [][]int
}

func (reporter *BasicReporter) Track(actual, ideal int) {
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

func (reporter *BasicReporter) Report() *Report {
	panic("implement me")
}
