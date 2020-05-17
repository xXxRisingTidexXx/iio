package observation

import (
	"gonum.org/v1/gonum/mat"
)

type Observer interface {
	Observe(float64)
	Expound() mat.Matrix
}
