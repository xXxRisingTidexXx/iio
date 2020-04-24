package guts

import (
	"gonum.org/v1/gonum/mat"
)

type Neuron interface {
	Evaluate(mat.Vector) mat.Vector
	Differentiate(mat.Vector) mat.Vector
}
