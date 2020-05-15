package networks

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/costs"
	"iio/pkg/initial"
	"iio/pkg/layered"
	"iio/pkg/loading"
	"iio/pkg/reports"
	"sync"
	"time"
)

func NewFeedForwardNetwork(
	epochNumber int,
	batchSize int,
	learningRate float64,
	trainingLoader loading.Loader,
	testLoader loading.Loader,
	weightInitializer initial.Initializer,
	biasInitializer initial.Initializer,
	costFunction costs.CostFunction,
	schemas ...*layered.Schema,
) *FeedForwardNetwork {
	if epochNumber < 1 {
		panic(fmt.Sprintf("networks: invalid epoch number, %d", epochNumber))
	}
	if batchSize < 1 {
		panic(fmt.Sprintf("networks: invalid batch size, %d", batchSize))
	}
	if trainingLoader == nil {
		panic("networks: feed forward network training loader can't be nil")
	}
	if testLoader == nil {
		panic("networks: feed forward network test loader can't be nil")
	}
	if weightInitializer == nil {
		panic("networks: feed forward network weight initializer can't be nil")
	}
	if biasInitializer == nil {
		panic("networks: feed forward network bias initializer can't be nil")
	}
	if costFunction == nil {
		panic("networks: feed forward network cost function can't be nil")
	}
	if schemas == nil {
		panic("networks: feed forward network schemas can't be nil")
	}
	length := len(schemas) - 1
	if length < 1 {
		panic(fmt.Sprintf("networks: invalid schema number (%d) and at least 2 required", length+1))
	}
	layers := make([]layered.Layer, length)
	for i, schema := range schemas {
		if schema == nil {
			panic(fmt.Sprintf("networks: nil schema at %d", i))
		}
		if i > 0 {
			if schema.Neuron == nil {
				panic(fmt.Sprintf("networks: input schema at %d", i))
			}
			layers[i-1] = layered.NewBasicLayer(
				schema.Neuron,
				weightInitializer.InitializeMatrix(schema.Size, schemas[i-1].Size),
				biasInitializer.InitializeVector(schema.Size),
			)
		} else if schema.Neuron != nil {
			panic("networks: the first schema must be an input one")
		}
	}
	return &FeedForwardNetwork{
		epochNumber,
		batchSize,
		learningRate,
		trainingLoader,
		testLoader,
		layers,
		costFunction,
		reports.NewBasicReporter(length),
	}
}

type FeedForwardNetwork struct {
	epochNumber    int
	batchSize      int
	learningRate   float64
	trainingLoader loading.Loader
	testLoader     loading.Loader
	layers         []layered.Layer
	costFunction   costs.CostFunction
	reporter       reports.Reporter
}

func (network *FeedForwardNetwork) Train() {
	for epoch := 0; epoch < network.epochNumber; epoch++ {
		start := time.Now()
		network.trainingLoader.Shuffle()
		for network.trainingLoader.Next() {
			batch := network.trainingLoader.Batch(network.batchSize)
			length := len(batch)
			deltasChannel := make(chan []*layered.Delta, length)
			waitGroup := &sync.WaitGroup{}
			waitGroup.Add(length)
			for _, sample := range batch {
				go network.train(sample, deltasChannel, waitGroup)
			}
			waitGroup.Wait()
			close(deltasChannel)
			learningRate := -network.learningRate / float64(length)
			for deltas := range deltasChannel {
				for i, layer := range network.layers {
					layer.Update(learningRate, deltas[i])
				}
			}
		}
		fmt.Printf("epoch %d took %s\n", epoch, time.Since(start))
	}
}

func (network *FeedForwardNetwork) train(
	sample *loading.Sample,
	deltasChannel chan<- []*layered.Delta,
	waitGroup *sync.WaitGroup,
) {
	length := len(network.layers)
	activations := make([]mat.Vector, length+1)
	activations[0] = sample.Data
	for i, layer := range network.layers {
		activations[i+1] = layer.FeedForward(activations[i])
	}
	deltas := make([]*layered.Delta, length)
	diffs := network.costFunction.Differentiate(activations[length], sample.Label)
	for i := length - 1; i >= 0; i-- {
		nodes := network.layers[i].ProduceNodes(diffs, activations[i+1])
		deltas[i] = layered.NewDelta(nodes, activations[i])
		if i > 0 {
			diffs = network.layers[i].BackPropagate(nodes)
		}
	}
	deltasChannel <- deltas
	waitGroup.Done()
}

func (network *FeedForwardNetwork) Test() *reports.Report {
	network.testLoader.Shuffle()
	for network.testLoader.Next() {
		batch := network.testLoader.Batch(network.batchSize)
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(len(batch))
		for _, sample := range batch {
			go network.test(sample, waitGroup)
		}
		waitGroup.Wait()
	}
	return network.reporter.Report()
}

func (network *FeedForwardNetwork) test(sample *loading.Sample, waitGroup *sync.WaitGroup) {
	activations := sample.Data
	for _, layer := range network.layers {
		activations = layer.FeedForward(activations)
	}
	network.reporter.Track(network.costFunction.Evaluate(activations), sample.Label)
	waitGroup.Done()
}

func (network *FeedForwardNetwork) Evaluate(input mat.Vector) int {
	panic("implement me")
}
