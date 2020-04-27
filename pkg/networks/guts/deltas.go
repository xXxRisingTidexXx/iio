package guts

import (
	"gonum.org/v1/gonum/mat"
)

func NewDeltas(nodes []mat.Vector, activations []mat.Vector) *Deltas {
	panic("implement me")
}

type Deltas struct {
	items []*Delta
}

func (deltas *Deltas) Add(other *Deltas) *Deltas {
	panic("implement me")
}

func (deltas *Deltas) Scale(alpha float64) *Deltas {
	panic("implement me")
}

func (deltas *Deltas) Get(i int) *Delta {
	panic("implement me")
}
