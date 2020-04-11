package layers

import (
	"iio/pkg/networks/neurons"
	"iio/pkg/vectors"
)

type ParallelLayer struct {
	neurons []neurons.Neuron
}

func (layer *ParallelLayer) Forward(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}
