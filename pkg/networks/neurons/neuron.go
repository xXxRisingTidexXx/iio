package neurons

import "iio/pkg/vectors"

// Main network's building block, which aggregates the input
// signal, estimates it and produces a single numeric output.
// Generally, the main difference between various neuron's
// implementations is the activation function, which performs
// overall operations. It's kind depends on the functionality
// we want to simulate.
type Neuron interface {
	// Bypasses the input signal to produce a scalar output.
	Forward(vectors.Vector) float64
}
