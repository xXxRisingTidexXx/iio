package guts

import (
	"iio/pkg/networks/guts/neurons"
	"iio/pkg/vectors"
)

type Layer interface {
	FeedForward(vectors.Vector) vectors.Vector
	BackPropagate(vectors.Vector) vectors.Vector
	Update([]*neurons.Bunch)
}
