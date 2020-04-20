package networks

import (
	"iio/pkg/sampling"
	"iio/pkg/vectors"
)

// Classification network generalization, which is in charge of
// image recognition.
type Network interface {
	// Functionality to "teach" the underlying layouts using a
	// pack of labeled objects. This process implies forward and
	// backward propagation, whose target is to adjusted the
	// underlying neurons' weights in an optimal way.
	Train([]*sampling.Sample)

	Validate([]*sampling.Sample)

	// Checks the accuracy of the net, penetrating it with a set
	// of new, more complicated samples. The returning value should
	// be treated as the integrated classifier's accuracy indicator.
	// Lies between 0 and 1 inclusively as well.
	Test([]*sampling.Sample) float64

	// Recognizes a single image, using the underlying configured
	// set of activation functions. Returns the recognized digit.
	Evaluate(vectors.Vector) byte
}
