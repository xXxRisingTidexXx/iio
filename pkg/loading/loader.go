package loading

import "iio/pkg/networks"

// High-level abstraction describing a data fetcher which
// downloads info from the specified data source and splits
// it into training and test sets.
type Loader interface {
	// Fetches a pack of labelled data and splits it into
	// two unique groups - training one and test one.
	Load() ([]*networks.Sample, []*networks.Sample, error)
}
