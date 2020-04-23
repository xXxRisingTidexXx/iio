package networks

import (
	"iio/pkg/networks/reports"
	"iio/pkg/sampling"
	"iio/pkg/vectors"
)

type Network interface {
	Train(*sampling.Samples)
	Validate(*sampling.Samples)
	Test(*sampling.Samples) reports.Report
	Evaluate(vectors.Vector) byte
}
