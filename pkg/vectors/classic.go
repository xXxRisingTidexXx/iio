package vectors

type ClassicVector struct {
	items []float64
}

func (vector *ClassicVector) Length() uint64 {
	return uint64(len(vector.items))
}

func (vector *ClassicVector) Get(uint64) float64 {
	panic("implement me")
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
