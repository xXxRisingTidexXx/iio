package loading

import (
	"iio/pkg/sampling"
)

// High-level abstraction describing a data fetcher which
// downloads info from the specified data source and splits
// it into training, validation and test sets.
type Loader interface {
	// Fetches a pack of labelled data and splits it into
	// three unique groups - training, validation test ones.
	Load() (*sampling.Samples, *sampling.Samples, *sampling.Samples)
}
