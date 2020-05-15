package main

import (
	"fmt"
	"iio/pkg/costs"
	"iio/pkg/initial"
	"iio/pkg/layered"
	"iio/pkg/loading"
	"iio/pkg/networks"
	"iio/pkg/neurons"
	"time"
)

func main() {
	trainingLoader, testLoader := loading.NewMNISTLoaders()
	network := networks.NewFeedForwardNetwork(
		5,
		32,
		0.01,
		trainingLoader,
		testLoader,
		initial.NewGlorotInitializer(),
		initial.NewZeroInitializer(),
		costs.NewMSECostFunction(),
		layered.NewInputSchema(784),
		layered.NewSchema(neurons.NewSigmoidNeuron(), 30),
		layered.NewSchema(neurons.NewSigmoidNeuron(), 10),
	)
	start := time.Now()
	network.Train()
	fmt.Printf("elapsed time: %s\n", time.Since(start))
}
