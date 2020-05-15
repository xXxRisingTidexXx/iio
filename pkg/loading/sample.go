package loading

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewSample(data mat.Vector, label int) *Sample {
	if data == nil {
		panic("loading: sample data can't be nil")
	}
	if label < 0 {
		panic("loading: class label must be non-negative")
	}
	return &Sample{data, label}
}

type Sample struct {
	Data  mat.Vector
	Label int
}

func (sample *Sample) Equal(other *Sample) bool {
	return other != nil && mat.Equal(sample.Data, other.Data) && sample.Label == other.Label
}

func (sample *Sample) String() string {
	return fmt.Sprintf("{%v %d}", sample.Data, sample.Label)
}
