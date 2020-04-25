package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/networks/guts"
	"iio/pkg/sampling"
	"sync"
)

type feedforwardNetwork struct {
	layers       []guts.Layer
	epochs       int
	batchSize    int
	learningRate float64
}

func (network *feedforwardNetwork) train(samples *sampling.Samples) {
	for epoch := 0; epoch < network.epochs; epoch++ {
		newSamples := samples.Shuffle()
		for newSamples.Next() {
			batch := newSamples.Batch(network.batchSize)
			length := batch.Length()
			waitGroup := &sync.WaitGroup{}
			waitGroup.Add(length)
			deltasChannel := make(chan *guts.Deltas, length)
			for i := 0; i < length; i++ {
				go network.propagate(batch.Get(i), waitGroup, deltasChannel)
			}
			waitGroup.Wait()
			totalDeltas := guts.NewDeltas(nil, nil)
			for deltas := range deltasChannel {
				totalDeltas = totalDeltas.Add(deltas)
			}
			totalDeltas = totalDeltas.Scale(-network.learningRate / float64(length))
			for i, layer := range network.layers {
				layer.Update(totalDeltas.Get(i))
			}
		}
	}
}

func (network *feedforwardNetwork) propagate(
	sample *sampling.Sample,
	waitGroup *sync.WaitGroup,
	deltasChannel chan<- *guts.Deltas,
) {
	defer waitGroup.Done()
	length := len(network.layers)
	activations := make([]mat.Vector, length+1)
	activations[0] = sample.Activations
	for i, layer := range network.layers {
		activations[i+1] = layer.FeedForward(activations[i])
	}
	nodes := make([]mat.Vector, length)

	for i := length - 2; i >= 0; i-- {
		
	}
	deltasChannel <- guts.NewDeltas(nodes, activations)
}

func (network *feedforwardNetwork) validate(samples *sampling.Samples) {
	panic("implement me")
}

func (network *feedforwardNetwork) test(samples *sampling.Samples) Report {
	panic("implement me")
}
