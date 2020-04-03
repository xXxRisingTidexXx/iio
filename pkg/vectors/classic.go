package vectors

type ClassicVector struct {
	items []float64
}

func (vector *ClassicVector) Length() int {
	return len(vector.items)
}

func (vector *ClassicVector) Get(i int) float64 {
	return vector.items[i]
}

func (vector *ClassicVector) Plus(Vector) Vector {
	panic("implement me")
}

func (vector *ClassicVector) Minus(Vector) Vector {
	panic("implement me")
}

func (vector *ClassicVector) TimesBy(float64) Vector {
	panic("implement me")
}

func (vector *ClassicVector) String() string {
	panic("implement me")
}
