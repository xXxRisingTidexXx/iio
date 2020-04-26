package guts

type Layers struct {
	items  []Layers
	length int
}

func (layers *Layers) Length() int {
	panic("implement me")
}

func (layers *Layers) Get(i int) Layer {
	panic("implement me")
}

func (layers *Layers) Last() Layer {
	panic("implement me")
}
