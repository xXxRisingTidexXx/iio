package networks

import (
	"iio/pkg/networks/layers"
	"iio/pkg/sampling"
)

type feedforwardNetwork struct {
	layers       []layers.Layer
	epochs       int
	batchSize    int
	learningRate float64
}

func (network *feedforwardNetwork) train(samples []*sampling.Sample) {
	for epoch := 0; epoch < network.epochs; epoch++ {

	}
}
