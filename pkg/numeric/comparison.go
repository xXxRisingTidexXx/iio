package numeric

import "math"

const AbsoluteTolerance = 1e-12

func IsClose(a, b float64) bool {
	return math.Abs(a-b) <= AbsoluteTolerance
}

func IsNull(a float64) bool {
	return IsClose(a, 0)
}
