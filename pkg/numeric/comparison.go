package numeric

import "math"

// Float number accuracy at the project boundaries.
const AbsoluteTolerance = 1e-12

// Main float number comparison method. It must be used in any
// computation or condition to avoid float value tricks and traps.
func IsClose(a, b float64) bool {
	return math.Abs(a-b) <= AbsoluteTolerance
}

// A shortcut for float comparison with zero.
func IsNull(a float64) bool {
	return IsClose(a, 0)
}
