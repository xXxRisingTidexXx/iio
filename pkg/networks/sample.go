package networks

import (
	"iio/pkg/vectors"
)

// An object representing labeled image suitable for a network
// classification.
type Sample struct {
	Activations vectors.Vector
	Label       byte
}
