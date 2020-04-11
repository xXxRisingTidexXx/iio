package layers

import "iio/pkg/vectors"

type Layer interface {
	Forward(vectors.Vector) vectors.Vector
}
