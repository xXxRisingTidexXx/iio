package networks

import (
	"fmt"
	"iio/pkg/costs"
	"iio/pkg/initial"
	"iio/pkg/layered"
	"iio/pkg/loading"
)

func NewOptions(
	epochNumber int,
	batchSize int,
	learningRate float64,
	trainingLoader loading.Loader,
	testLoader loading.Loader,
	weightInitializer initial.Initializer,
	biasInitializer initial.Initializer,
	costFunction costs.CostFunction,
	schemas ...*layered.Schema,
) *Options {
	if epochNumber < 1 {
		panic(fmt.Sprintf("networks: invalid epoch number, %d", epochNumber))
	}
	if batchSize < 1 {
		panic(fmt.Sprintf("networks: invalid batch size, %d", batchSize))
	}
	if trainingLoader == nil {
		panic("networks: network training loader can't be nil")
	}
	if length := trainingLoader.Length(); length < 1 {
		panic(fmt.Sprintf("networks: network training set has invalid length, %d", length))
	}
	if testLoader == nil {
		panic("networks: network test loader can't be nil")
	}
	if length := testLoader.Length(); length < 1 {
		panic(fmt.Sprintf("networks: network test set has invalid length, %d", length))
	}
	if weightInitializer == nil {
		panic("networks: network weight initializer can't be nil")
	}
	if biasInitializer == nil {
		panic("networks: network bias initializer can't be nil")
	}
	if costFunction == nil {
		panic("networks: network cost function can't be nil")
	}
	if schemas == nil {
		panic("networks: network schemas can't be nil")
	}
	if length := len(schemas); length < 2 {
		panic(fmt.Sprintf("networks: invalid schema number (%d) and at least 2 required", length))
	}
	for i, schema := range schemas {
		if schema == nil {
			panic(fmt.Sprintf("networks: nil schema at %d", i))
		}
		if i > 0 && schema.Neuron == nil {
			panic(fmt.Sprintf("networks: input schema at %d", i))
		} else if schema.Neuron != nil {
			panic("networks: the first schema must be an input one")
		}
	}
	return &Options{
		epochNumber,
		batchSize,
		learningRate,
		trainingLoader,
		testLoader,
		weightInitializer,
		biasInitializer,
		costFunction,
		schemas,
	}
}

type Options struct {
	EpochNumber       int
	BatchSize         int
	LearningRate      float64
	TrainingLoader    loading.Loader
	TestLoader        loading.Loader
	WeightInitializer initial.Initializer
	BiasInitializer   initial.Initializer
	CostFunction      costs.CostFunction
	Schemas           []*layered.Schema
}
