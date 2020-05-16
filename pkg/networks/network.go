package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/estimate"
)

type Network interface {
	Evaluate(mat.Vector) int
	Train()
	Test() *estimate.Report
}
