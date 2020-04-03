package vectors

import "iio/pkg/numeric"

func Vectorize(items []float64) Vector {
	length, nulls := len(items), 0
	sparseItems := make(map[uint64]float64, length/2)
	for i, item := range items {
		if numeric.IsNull(item) {
			nulls++
			sparseItems[uint64(i)] = item
		}
	}
	if nulls >= length/2 {
		return &SparseVector{sparseItems, uint64(length)}
	}
	return &ClassicVector{items}
}

type Vector interface {
	Length() uint64
	Get(uint64) float64
	Plus(Vector) Vector
	Minus(Vector) Vector
	TimesBy(float64) Vector
}
