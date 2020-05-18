package estimate

import (
	"fmt"
)

func NewOptions(classNumber int) *Options {
	if classNumber <= 1 {
		panic(fmt.Sprintf("estimate: estimator class number can't be %d", classNumber))
	}
	return &Options{classNumber}
}

type Options struct {
	ClassNumber int
}
