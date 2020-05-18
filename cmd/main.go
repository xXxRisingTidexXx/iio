package main

import (
	"fmt"
	"iio/pkg/costs"
	"iio/pkg/initial"
	"iio/pkg/layered"
	"iio/pkg/loading"
	"iio/pkg/networks"
	"time"
)

func main() {
	trainingLoader, testLoader := loading.NewMNISTLoaders()
	network := networks.NewFeedForwardNetwork(
		networks.NewOptions(
			1,
			16,
			0.01,
			trainingLoader,
			testLoader,
			initial.NewGlorotInitializer(),
			initial.NewZeroInitializer(),
			costs.NewMSECostFunction(),
			layered.NewInputSchema(784),
			layered.NewSigmoidSchema(30),
			layered.NewSigmoidSchema(10),
		),
	)
	start := time.Now()
	network.Train()
	fmt.Printf("training elapsed time: %s\n", time.Since(start))
	start = time.Now()
	report := network.Test()
	fmt.Printf("test elapsed time: %s\n\n", time.Since(start))
	fmt.Println(report)
}
