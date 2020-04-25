package sampling

import (
	"gonum.org/v1/gonum/mat"
)

type Sample struct {
	Activations mat.Vector
	Label       int
}
