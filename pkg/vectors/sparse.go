package vectors

import (
	"fmt"
)

// Optimized linear data structure holding a map of non-zero
// float64 values. The main benefit lies in the general vector
// size: arrays with many (roughly less than a half) zero cells
// can be tighten, then the general memory usage can be
// decreased. Probably, this vector would be less useful in the
// case of dense data (in that way use vector.ClassicVector ).
type SparseVector struct {
	items  map[int]float64
	length int
}

func (vector *SparseVector) Length() int {
	return vector.length
}

func (vector *SparseVector) Get(i int) float64 {
	if i < 0 || i >= vector.length {
		panic(fmt.Sprintf("out of bounds: length %d, index %d", vector.length, i))
	}
	if item, ok := vector.items[i]; ok {
		return item
	}
	return 0.0
}

func (vector *SparseVector) Plus(Vector) Vector {
	panic("implement me")
}

func (vector *SparseVector) Minus(Vector) Vector {
	panic("implement me")
}

func (vector *SparseVector) TimesBy(float64) Vector {
	panic("implement me")
}

func (vector *SparseVector) String() string {
	return Shorten(vector)
}
