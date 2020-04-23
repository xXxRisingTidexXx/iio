package vectors

import (
	"math"
)

// Float number accuracy at the project boundaries.
const absoluteTolerance = 1e-12

// Main float number comparison method. It must be used in any
// computation or condition to avoid float value tricks and traps.
func isClose(a, b float64) bool {
	return math.Abs(a-b) <= absoluteTolerance
}

// A shortcut for float comparison with zero.
func isNull(a float64) bool {
	return isClose(a, 0)
}
