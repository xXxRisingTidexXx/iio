package guts

import (
	"iio/pkg/networks/guts/neurons"
	"iio/pkg/vectors"
)

type feedforwardLayer struct {
	neurons []Neuron
}

func (layer *feedforwardLayer) feedForward(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}

func (layer *feedforwardLayer) backPropagate(deltas vectors.Vector) vectors.Vector {
	panic("implement me")
}

func (layer *feedforwardLayer) update(bunches []*neurons.Bunch) {
	panic("implement me")
}
