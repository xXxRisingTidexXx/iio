package vectors

import (
	"fmt"
	"strings"
)

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
	builder := strings.Builder{}
	builder.WriteString("[ ")
	if vector.length <= 10 {
		for i := 0; i < vector.length; i++ {
			builder.WriteString(fmt.Sprintf("%.3f ", vector.Get(i)))
		}
	} else {
		for i := 0; i < 3; i++ {
			builder.WriteString(fmt.Sprintf("%.3f ", vector.Get(i)))
		}
		builder.WriteString("... ")
		for i := vector.length - 3; i < vector.length; i++ {
			builder.WriteString(fmt.Sprintf("%.3f ", vector.Get(i)))
		}
	}
	builder.WriteString("]")
	return builder.String()
}
