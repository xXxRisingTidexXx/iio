package layers

import (
	"iio/pkg/networks/layers/neurons"
	"iio/pkg/vectors"
)

type OutputLayer struct {
	layer *feedforwardLayer
}

func (layer *HiddenLayer) FeedForward(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}

func (layer *HiddenLayer) BackPropagate(activations vectors.Vector) vectors.Vector {
	panic("implement me")
}

func (layer *HiddenLayer) Update(bunches []*neurons.Bunch) {
	panic("implement me")
}