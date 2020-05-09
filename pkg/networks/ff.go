package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/guts"
	"sync"
)

type FFNetwork struct {
	epochs       int
	batchSize    int
	learningRate float64
	layers       *guts.Layers
	costFunction guts.CostFunction
}

func (network *FFNetwork) Train(samples *sampling.Samples) {
	for epoch := 0; epoch < network.epochs; epoch++ {
		newSamples := samples.Shuffle()
		for newSamples.Next() {
			batch := newSamples.Batch(network.batchSize)
			length := batch.Length()
			waitGroup := &sync.WaitGroup{}
			waitGroup.Add(length)
			deltasChannel := make(chan *guts.Deltas, length)
			for i := 0; i < length; i++ {
				go network.train(batch.Get(i), waitGroup, deltasChannel)
			}
			waitGroup.Wait()
			totalDeltas := guts.NewDeltas(nil, nil)
			for deltas := range deltasChannel {
				totalDeltas = totalDeltas.Add(deltas)
			}
			totalDeltas = totalDeltas.Scale(-network.learningRate / float64(length))
			for i := 0; i < network.layers.Length(); i++ {
				network.layers.Get(i).Update(totalDeltas.Get(i))
			}
		}
	}
}

func (network *FFNetwork) train(
	sample *sampling.Sample,
	waitGroup *sync.WaitGroup,
	deltasChannel chan<- *guts.Deltas,
) {
	defer waitGroup.Done()
	length := network.layers.Length()
	activations := make([]mat.Vector, length+1)
	activations[0] = sample.Activations()
	for i := 0; i < length; i++ {
		activations[i+1] = network.layers.Get(i).FeedForward(activations[i])
	}
	nodes := make([]mat.Vector, length)
	nodes[length-1] = network.layers.Last().ProduceNodes(
		network.costFunction.Evaluate(activations[length], sample.Label()),
	)
	for i := length - 2; i >= 0; i-- {
		nodes[i] = network.layers.Get(i).ProduceNodes(
			network.layers.Get(i + 1).BackPropagate(nodes[i+1]),
		)
	}
	deltasChannel <- guts.NewDeltas(nodes, activations)
}

func (network *FFNetwork) Validate(samples *sampling.Samples) {
	panic("implement me")
}

func (network *FFNetwork) Test(samples *sampling.Samples) Report {
	panic("implement me")
}

func (network *FFNetwork) Evaluate(activations mat.Vector) int {
	panic("implement me")
}
