package sampling

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewSample(activations mat.Vector, label int) *Sample {
	if activations == nil {
		panic(fmt.Sprintf("sampling: sample activations can't be nil"))
	}
	return &Sample{activations, label}
}

type Sample struct {
	activations mat.Vector
	label       int
}

func (sample *Sample) Activations() mat.Vector {
	return sample.activations
}

func (sample *Sample) Label() int {
	return sample.label
}

func (sample *Sample) Equal(other *Sample) bool {
	return sample == other ||
		other != nil &&
			mat.Equal(sample.activations, other.activations) &&
			sample.label == other.label
}
