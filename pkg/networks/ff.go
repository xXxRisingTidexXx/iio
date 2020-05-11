package networks

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/guts"
	"iio/pkg/loading"
)

type FFNetwork struct {
	layers         []guts.Layer
	costFunction   guts.CostFunction
	trainingLoader loading.Loader
	testLoader     loading.Loader
	epochs         int
	batchSize      int
	learningRate   float64
}

func (network *FFNetwork) Train() {
	for epoch := 0; epoch < network.epochs; epoch++ {
		network.trainingLoader.Shuffle()
		for network.trainingLoader.Next() {
			batch := network.trainingLoader.Batch(network.batchSize)
			length := len(batch)
			deltasChannel := make(chan []*guts.Delta, length)
			for _, sample := range batch {
				go network.train(sample, deltasChannel)
			}
			learningRate := -network.learningRate / float64(length)
			for deltas := range deltasChannel {
				for i, layer := range network.layers {
					layer.Update(learningRate, deltas[i])
				}
			}
		}
	}
}

func (network *FFNetwork) train(sample *loading.Sample, deltasChannel chan<- []*guts.Delta) {
	length := len(network.layers)
	activations := make([]mat.Vector, length+1)
	activations[0] = sample.Data()
	for i, layer := range network.layers {
		activations[i+1] = layer.FeedForward(activations[i])
	}
	deltas := make([]*guts.Delta, length)
	diffs := network.costFunction.Differentiate(activations[length], sample.Label())
	for i := length - 1; i >= 0; i-- {
		nodes := network.layers[i].ProduceNodes(diffs, activations[i+1])
		deltas[i] = guts.NewDelta(nodes, activations[i])
		if i > 0 {
			diffs = network.layers[i].BackPropagate(nodes)
		}
	}
	deltasChannel <- deltas
}

func (network *FFNetwork) Test() *Report {
	network.testLoader.Shuffle()
	for network.testLoader.Next() {
		batch := network.testLoader.Batch(network.batchSize)
		resultChannel := make(chan *result, len(batch))
		for _, sample := range batch {
			go network.test(sample, resultChannel)
		}
		for result := range resultChannel {
			fmt.Println(result)
			// Do some logic with results and report
		}
	}
	return &Report{}
}

func (network *FFNetwork) test(sample *loading.Sample, resultChannel chan<- *result) {
	panic("implement me")
}

func (network *FFNetwork) Evaluate(input mat.Vector) int {
	panic("implement me")
}
