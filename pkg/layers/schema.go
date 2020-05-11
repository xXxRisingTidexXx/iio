package layers

import (
	"fmt"
	"iio/pkg/neurons"
)

func NewSchema(neuron neurons.Neuron, size int) *Schema {
	if size < 1 {
		panic(fmt.Sprintf("layers: schema got invalid size, %d", size))
	}
	return &Schema{neuron, size}
}

type Schema struct {
	Neuron neurons.Neuron
	Size   int
}
