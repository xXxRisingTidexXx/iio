package layers

import "iio/pkg/vectors"

// Complex abstraction managing a set of activation functions
// (neurons). It coordinates an activation signal computation,
// propagating data forwards straight to the next layer. Indeed,
// this entity should recognize some patterns at samples and
// report this important information to its followers to integrate
// separate breadcrumbs into standalone contours.
type Layer interface {
	// Passes the input signal, transforms it and produces some
	// vector output. The amount of underlying computation units
	// is arbitrary, but each neuron should have the same number
	// of inputs as the previous layer neuron quantity.
	Forward(vectors.Vector) vectors.Vector
}
