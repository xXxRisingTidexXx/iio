package vectors

import (
	"fmt"
	"strings"
)

// A simple and elegant method to stringify a numeric array.
// Too long parts are cut and replaced by three dots.
func Shorten(vector Vector) string {
	builder := strings.Builder{}
	builder.WriteString("[ ")
	if vector.Length() <= 20 {
		for i := 0; i < vector.Length(); i++ {
			builder.WriteString(fmt.Sprintf("%.3f ", vector.Get(i)))
		}
	} else {
		for i := 0; i < 6; i++ {
			builder.WriteString(fmt.Sprintf("%.3f ", vector.Get(i)))
		}
		builder.WriteString("... ")
		for i := vector.Length() - 6; i < vector.Length(); i++ {
			builder.WriteString(fmt.Sprintf("%.3f ", vector.Get(i)))
		}
	}
	builder.WriteString("]")
	return builder.String()
}
