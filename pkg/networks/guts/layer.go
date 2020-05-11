package guts

import (
	"gonum.org/v1/gonum/mat"
)

type Layer interface {
	FeedForward(activations mat.Vector) mat.Vector
	ProduceNodes(diffs, activations mat.Vector) mat.Vector
	BackPropagate(nodes mat.Vector) mat.Vector
	Update(delta *Delta)
}
