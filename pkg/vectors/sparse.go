package vectors

import "fmt"

type SparseVector struct {
	items  map[uint64]float64
	length uint64
}

func (vector *SparseVector) Length() uint64 {
	return vector.length
}

func (vector *SparseVector) Get(uint64) float64 {
	panic("implement me")
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
