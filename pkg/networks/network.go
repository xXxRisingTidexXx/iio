package networks

import (
	"iio/pkg/sampling"
	"iio/pkg/vectors"
)

type Network interface {
	Train(*sampling.Samples)
	Validate(*sampling.Samples)
	Test(*sampling.Samples) Report
	Evaluate(vectors.Vector) byte
}
