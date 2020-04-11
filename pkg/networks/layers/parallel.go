package layers

import (
	"iio/pkg/networks/neurons"
	"iio/pkg/vectors"
)

// Specific neuron implementation, which leverages the key golang
// feature - goroutines, channels ad esy-to-use concurrency.
// Actually, the parallelization lies in simultaneous neuron
// computation.
type ParallelLayer struct {
	neurons []neurons.Neuron
}

func (layer *ParallelLayer) Forward(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}
