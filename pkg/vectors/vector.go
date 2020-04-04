package vectors

import "iio/pkg/numeric"

// Transforms "raw" numerical slice into specific vector.
// Implementation type depends on the slice content: if it
// contains too many zero values, the future vector will be
// sparse; otherwise, classic slice-based vector will be used.
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

// Main numeric 1D-array interface. Declares a set of methods
// for convenient mathematical computations like search by
// index, addition, subtraction, constant multiplication. In
// the nearest future there gonna appear dot and cross products,
// inverse matrix computation, etc.
type Vector interface {
	Length() int
	Get(int) float64
	Plus(Vector) Vector
	Minus(Vector) Vector
	TimesBy(float64) Vector
}
