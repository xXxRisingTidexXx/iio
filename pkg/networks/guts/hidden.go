package guts

import (
	"iio/pkg/networks/guts/neurons"
	"iio/pkg/vectors"
)

type HiddenLayer struct {
	layer *feedforwardLayer
}

func (layer *OutputLayer) FeedForward(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}

func (layer *OutputLayer) BackPropagate(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}

func (layer *OutputLayer) Update(bunches []*neurons.Bunch) {
	panic("implement me")
}