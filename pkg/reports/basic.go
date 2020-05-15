package reports

import (
	"fmt"
)

func NewBasicReporter(classNumber int) *BasicReporter {
	if classNumber <= 1 {
		panic(fmt.Sprintf("reports: got invalid class number, %d", classNumber))
	}
	return &BasicReporter{}
}

type BasicReporter struct {

}

func (reporter *BasicReporter) Track(actual, ideal int) {
	panic("implement me")
}

func (reporter *BasicReporter) Report() *Report {
	panic("implement me")
}

