package loading

import (
	"iio/pkg/vectors"
)

// An object representing labeled image suitable for a network
// classification.
type Example struct {
	Image vectors.Vector
	Label byte
}
