package observation

import (
	"fmt"
)

func NewOptions(epochNumber, setLength, portionSize int) *Options {
	if epochNumber < 1 {
		panic(fmt.Sprintf("observer: basic observer got invalid epoch number, %d", epochNumber))
	}
	if setLength < 1 {
		panic(fmt.Sprintf("observer: basic observer got invalid set length, %d", setLength))
	}
	if portionSize < 1 {
		panic(fmt.Sprintf("observer: basic observer got invalid portion size, %d", portionSize))
	}
	return &Options{epochNumber, setLength, portionSize}
}

type Options struct {
	EpochNumber int
	SetLength   int
	PortionSize int
}
