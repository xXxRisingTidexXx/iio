package sampling

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Common example constructor. Actually, this struct needs no checks
// except of the vector nil assertion.
func NewSample(activations mat.Vector, label int) *Sample {
	if activations == nil {
		panic(fmt.Sprintf("sampling: sample activations can't be nil"))
	}
	return &Sample{activations, label}
}

// Simple abstraction encapsulating learning instance data. It
// contains example's vector view and integer class label suitable
// for a network processing. The same activations must have the same
// labels but the same labels may refer to different vectors.
type Sample struct {
	// Learning object vector representation. For instance, 2D images
	// should be flattened into 1D arrays to satisfy this struct's
	// contract.
	activations mat.Vector

	// Integer class mapping.
	label int
}

// Public vector accessor.
func (sample *Sample) Activations() mat.Vector {
	return sample.activations
}

// Public class mapping accessor.
func (sample *Sample) Label() int {
	return sample.label
}

func (sample *Sample) Equal(other *Sample) bool {
	return sample == other ||
		other != nil &&
			mat.Equal(sample.activations, other.activations) &&
			sample.label == other.label
}

func (sample *Sample) String() string {
	return fmt.Sprintf("{%v %d}", sample.activations, sample.label)
}
