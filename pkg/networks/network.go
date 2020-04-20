package networks

import (
	"iio/pkg/networks/reports"
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

	// Can be used to check the net hyperparameters.
	Validate([]*sampling.Sample)

	// Checks the accuracy of the net, penetrating it with a set
	// of new, more complicated samples. Returns the overall
	// conclusion of the check with common metrics.
	Test([]*sampling.Sample) reports.Report

	// Recognizes a single image, using the underlying configured
	// set of activation functions. Returns the recognized digit.
	Evaluate(vectors.Vector) byte
}
