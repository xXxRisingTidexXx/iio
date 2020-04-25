package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
)

type Network interface {
	Train(*sampling.Samples)
	Validate(*sampling.Samples)
	Test(*sampling.Samples) Report
	Evaluate(mat.Vector) int
}
