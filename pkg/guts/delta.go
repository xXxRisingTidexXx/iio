package guts

import (
	"gonum.org/v1/gonum/mat"
)

type Delta struct {
	Weights mat.Matrix
	Biases  mat.Vector
}

func (delta *Delta) Scale(alpha float64) *Delta {
	panic("implement me")
}
