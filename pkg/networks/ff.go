package networks

import (
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
			for deltas := range deltasChannel {
				for i, layer := range network.layers {
					layer.Update(-network.learningRate/float64(length), deltas[i])
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
	nodes := make([]mat.Vector, length)
	nodes[length-1] = network.layers[length-1].ProduceNodes(
		network.costFunction.Evaluate(activations[length], sample.Label()),
	)
	for i := length - 2; i >= 0; i-- {
		nodes[i] = network.layers[i].ProduceNodes(network.layers[i+1].BackPropagate(nodes[i+1]))
		// Add here deltas calculation
	}
	deltasChannel <- []*guts.Delta{}
}

func (network *FFNetwork) Test() Report {
	panic("implement me")
}

func (network *FFNetwork) Evaluate(input mat.Vector) int {
	panic("implement me")
}
