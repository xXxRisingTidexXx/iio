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
	// Calculates the array element amount.
	Length() int

	// Selects the element at the specified index. Should panic
	// in case of a missing index or out-of-bounds-error.
	Get(int) float64

	// Sums two independent arrays producing the third independent
	// one. Should panic in case of different array length.
	Plus(Vector) Vector

	// Subtracts element-wisely two vectors of the same length. In
	// case of different lengths should produce a panic.
	Minus(Vector) Vector

	// Performs element-wise multiplication of two vectors.
	// Different array length causes a panic.
	TimesBy(float64) Vector
}
