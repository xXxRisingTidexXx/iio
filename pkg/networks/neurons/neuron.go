package neurons

import "iio/pkg/vectors"

type Neuron interface {
	Forward(vectors.Vector) float64
}
