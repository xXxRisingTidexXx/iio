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
	Activations mat.Vector
	Label       int
}

//func (sample *Sample) Equal(other *Sample) bool {
//	if sample == other {
//		return true
//	}
//	if other == nil {
//		return false
//	}
//
//}
