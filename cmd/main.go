package main

import (
	"iio/pkg/costs"
	"iio/pkg/init"
	"iio/pkg/layered"
	"iio/pkg/loading"
	"iio/pkg/networks"
	"iio/pkg/neurons"
)

func main() {
	trainingLoader, testLoader := loading.NewMNISTLoaders()
	network := networks.NewFeedForwardNetwork(
		10,
		4,
		0.3,
		trainingLoader,
		testLoader,
		init.NewGlorotInitializer(),
		init.NewZeroInitializer(),
		costs.NewMSECostFunction(),
		layered.NewInputSchema(784),
		layered.NewSchema(neurons.NewSigmoidNeuron(), 30),
		layered.NewSchema(neurons.NewSigmoidNeuron(), 10),
	)
	network.Train()
}
