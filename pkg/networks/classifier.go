package networks

import "iio/pkg/vectors"

// Classification network generalization, which provides a set
// of useful methods for training, testing, direct classification
// and dump - network's weight storing to the specified file.
type Classifier interface {
	Train([]*Sample) float64
	Test([]*Sample) float64
	Classify(vectors.Vector) byte
	Dump(string) error
}
