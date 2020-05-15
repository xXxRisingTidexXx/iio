package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/estimate"
)

type Network interface {
	Train()
	Test() *estimate.Report
	Evaluate(mat.Vector) int
}
