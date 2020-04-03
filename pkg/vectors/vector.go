package vectors

import "iio/pkg/numeric"

func Vectorize(items []float64) Vector {
	length, nulls := len(items), 0
	sparseItems := make(map[int]float64, length/2)
	for i, item := range items {
		if numeric.IsNull(item) {
			nulls++
			sparseItems[i] = item
		}
	}
	if nulls >= length/2 {
		return &SparseVector{sparseItems, length}
	}
	return &ClassicVector{items}
}

type Vector interface {
	Length() int
	Get(int) float64
	Plus(Vector) Vector
	Minus(Vector) Vector
	TimesBy(float64) Vector
}
