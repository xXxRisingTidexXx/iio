package vectors

import "fmt"

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
	return fmt.Sprintf("(%d,)\n%v\n", vector.length, vector.items)
}
