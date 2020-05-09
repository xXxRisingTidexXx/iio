package networks

import (
	"gonum.org/v1/gonum/mat"
)

type Network interface {
	Train()
	Test() Report
	Evaluate(mat.Vector) int
}
