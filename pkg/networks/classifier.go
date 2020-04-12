package networks

import "iio/pkg/vectors"

// Classification network generalization, which is in charge of
// image recognition.
type Classifier interface {
	// Functionality to "teach" the underlying layouts using a
	// pack of labeled objects. This process implies forward and
	// backward propagation, whose target is to adjusted the
	// underlying neurons' weights in an optimal way. Returns
	// the percent of successfully classified images (a number
	// from 0 to 1 inclusively).
	Train([]*Sample) float64

	// Checks the accuracy of the net, penetrating it with a set
	// of new, more complicated samples. The returning value should
	// be treated as the integrated classifier's accuracy indicator.
	// Lies between 0 and 1 inclusively as well.
	Test([]*Sample) float64

	// Recognizes a single image, using the underlying configured
	// set of activation functions. Returns the recognized digit.
	Classify(vectors.Vector) byte

	// Converts all inner digital infrastructure into specific
	// format to implement persistence - networks should "live"
	// out of RAM :) Accepts a destination path and indicates an
	// error if something goes wrong.
	Dump(string) error
}
