package guts

import (
	"gonum.org/v1/gonum/mat"
)

type Delta struct {
	Weights mat.Matrix
	Biases  mat.Vector
}
