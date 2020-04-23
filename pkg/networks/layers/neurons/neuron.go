package neurons

import (
	"iio/pkg/vectors"
)

type Neuron interface {
	FeedForward(vectors.Vector) float64
	Derivative(float64) float64
	BackPropagate(float64) vectors.Vector
	Update(*Bunch)
}
