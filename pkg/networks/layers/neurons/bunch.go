package neurons

import (
	"iio/pkg/vectors"
)

type Bunch struct {
	Weights vectors.Vector
	Bias    float64
}

func (bunch *Bunch) Plus(other *Bunch) *Bunch {
	panic("implement me")
}

func (bunch *Bunch) TimesBy(number float64) *Bunch {
	panic("implement me")
}
