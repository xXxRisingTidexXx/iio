package layers

import (
	"gonum.org/v1/gonum/mat"
)

type Layer interface {
	FeedForward(mat.Vector) mat.Vector
	ProduceNodes(mat.Vector, mat.Vector) mat.Vector
	BackPropagate(mat.Vector) mat.Vector
	Update(float64, *Delta)
}
