package guts

import "gonum.org/v1/gonum/mat"

type HiddenLayer struct {
	layer *feedforwardLayer
}

func (layer *HiddenLayer) FeedForward(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *HiddenLayer) BackPropagate(activations mat.Vector) mat.Vector {
	panic("implement me")
}

func (layer *HiddenLayer) Update(weightDeltas mat.Matrix, biasDeltas mat.Vector) {
	panic("implement me")
}
