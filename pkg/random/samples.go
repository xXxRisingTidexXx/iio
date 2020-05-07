package random

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
	"math/rand"
)

// Generates a random samples of a given length and item size.
// TODO: should be tested.
func NewSamples(
	length int,
	size int,
	minActivation float64,
	maxActivation float64,
	minLabel int,
	maxLabel int,
) *sampling.Samples {
	items := make([]*sampling.Sample, length)
	activationOffset, labelOffset := maxActivation-minActivation, maxLabel-minLabel
	for i := 0; i < length; i++ {
		floats := make([]float64, size)
		for j := 0; j < size; j++ {
			floats[j] = minActivation + rand.Float64()*activationOffset
		}
		items[i] = sampling.NewSample(mat.NewVecDense(size, floats), minLabel+rand.Intn(labelOffset))
	}
	return sampling.NewSamples(items...)
}
