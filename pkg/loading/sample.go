package loading

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewSample(data mat.Vector, label int) *Sample {
	if data == nil {
		panic(fmt.Sprintf("loading: sample data can't be nil"))
	}
	return &Sample{data, label}
}

type Sample struct {
	data  mat.Vector
	label int
}

func (sample *Sample) Data() mat.Vector {
	return sample.data
}

func (sample *Sample) Label() int {
	return sample.label
}

func (sample *Sample) Equal(other *Sample) bool {
	return sample == other || other != nil && mat.Equal(sample.data, other.data) && sample.label == other.label
}

func (sample *Sample) String() string {
	return fmt.Sprintf("{%v %d}", sample.data, sample.label)
}
