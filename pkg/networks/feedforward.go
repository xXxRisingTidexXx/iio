package networks

import (
	"iio/pkg/networks/guts"
	"iio/pkg/sampling"
)

type feedforwardNetwork struct {
	layers       []guts.Layer
	epochs       int
	batchSize    int
	learningRate float64
}

func (network *feedforwardNetwork) train(samples *sampling.Samples) {
	for epoch := 0; epoch < network.epochs; epoch++ {
		shuffled := samples.Shuffle()

	}
}

func (network *feedforwardNetwork) validate(samples *sampling.Samples) {
	panic("implement me")
}

func (network *feedforwardNetwork) test(samples *sampling.Samples) Report {
	panic("implement me")
}
