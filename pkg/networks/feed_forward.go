package networks

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/costs"
	"iio/pkg/estimate"
	"iio/pkg/initial"
	"iio/pkg/layered"
	"iio/pkg/loading"
	"iio/pkg/observation"
	"sync"
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
) Network {
	length := len(schemas) - 1
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
	return &feedForwardNetwork{
		epochNumber,
		batchSize,
		learningRate,
		trainingLoader,
		testLoader,
		layers,
		costFunction,
		observation.NewBasicObserver(epochNumber, trainingLoader.Length(), batchSize),
		estimate.NewBasicEstimator(schemas[length].Size),
	}
}

type feedForwardNetwork struct {
	epochNumber    int
	batchSize      int
	learningRate   float64
	trainingLoader loading.Loader
	testLoader     loading.Loader
	layers         []layered.Layer
	costFunction   costs.CostFunction
	observer       observation.Observer
	estimator      estimate.Estimator
}

func (network *feedForwardNetwork) Evaluate(input mat.Vector) int {
	panic("implement me")
}

func (network *feedForwardNetwork) Train() mat.Matrix {
	for epoch := 0; epoch < network.epochNumber; epoch++ {
		network.trainingLoader.Shuffle()
		for network.trainingLoader.Next() {
			batch := network.trainingLoader.Batch(network.batchSize)
			length := len(batch)
			learningRate := -network.learningRate / float64(length)
			deltasChannel := make(chan []*layered.Delta, length)
			waitGroup := &sync.WaitGroup{}
			waitGroup.Add(length)
			for _, sample := range batch {
				go network.train(sample, learningRate, deltasChannel, waitGroup)
			}
			waitGroup.Wait()
			close(deltasChannel)
			for deltas := range deltasChannel {
				waitGroup.Add(len(network.layers))
				for i, layer := range network.layers {
					go network.update(layer, deltas[i], waitGroup)
				}
				waitGroup.Wait()
			}
		}
	}
	return network.observer.Expound()
}

func (network *feedForwardNetwork) train(
	sample *loading.Sample,
	learningRate float64,
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
		deltas[i] = layered.NewDelta(nodes, activations[i], learningRate)
		if i > 0 {
			diffs = network.layers[i].BackPropagate(nodes)
		}
	}
	network.observer.Observe(network.costFunction.Cost(activations[length], sample.Label))
	deltasChannel <- deltas
	waitGroup.Done()
}

func (network *feedForwardNetwork) update(layer layered.Layer, delta *layered.Delta, waitGroup *sync.WaitGroup) {
	layer.Update(delta)
	waitGroup.Done()
}

func (network *feedForwardNetwork) Test() *estimate.Report {
	network.testLoader.Shuffle()
	for network.testLoader.Next() {
		batch := network.testLoader.Batch(network.batchSize)
		length := len(batch)
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(length)
		for _, sample := range batch {
			go network.test(sample, waitGroup)
		}
		waitGroup.Wait()
	}
	return network.estimator.Estimate()
}

func (network *feedForwardNetwork) test(sample *loading.Sample, waitGroup *sync.WaitGroup) {
	activations := sample.Data
	for _, layer := range network.layers {
		activations = layer.FeedForward(activations)
	}
	network.estimator.Track(network.costFunction.Evaluate(activations), sample.Label)
	waitGroup.Done()
}
