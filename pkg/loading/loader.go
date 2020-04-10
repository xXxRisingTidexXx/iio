package loading

import "iio/pkg/networks"

// High-level abstraction describing a data fetcher which
// downloads info from the specified data source and splits
// it into training and test sets.
type Loader interface {
	Load() ([]*networks.Sample, []*networks.Sample, error)
}
