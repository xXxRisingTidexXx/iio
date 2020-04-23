package layers

import (
	"iio/pkg/vectors"
)

type Layer interface {
	FeedForward(vectors.Vector) vectors.Vector
	BackPropagate(vectors.Vector) vectors.Vector
	Update()
}
