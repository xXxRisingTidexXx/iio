package networks

import (
	"gonum.org/v1/gonum/mat"
	"iio/pkg/reports"
)

type Network interface {
	Train()
	Test() *reports.Report
	Evaluate(mat.Vector) int
}
