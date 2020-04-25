package guts

import (
	"gonum.org/v1/gonum/mat"
)

type Layer interface {
	FeedForward(mat.Vector) mat.Vector
	BackPropagate(mat.Vector) mat.Vector
	Update(*Delta)
}
