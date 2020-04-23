package layers

import (
	"iio/pkg/networks/layers/neurons"
	"iio/pkg/vectors"
)

type Layer interface {
	FeedForward(vectors.Vector) vectors.Vector
	BackPropagate(vectors.Vector) vectors.Vector
	Update([]*neurons.Bunch)
}
