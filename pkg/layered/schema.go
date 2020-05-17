package layered

import (
	"fmt"
	"iio/pkg/neurons"
)

func NewInputSchema(size int) *Schema {
	return NewSchema(nil, size)
}

func NewSigmoidSchema(size int) *Schema {
	return NewSchema(neurons.NewSigmoidNeuron(), size)
}

func NewSchema(neuron neurons.Neuron, size int) *Schema {
	if size <= 1 {
		panic(fmt.Sprintf("layers: schema got invalid size, %d", size))
	}
	return &Schema{neuron, size}
}

type Schema struct {
	Neuron neurons.Neuron
	Size   int
}
