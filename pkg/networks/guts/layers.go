package guts

type Layers struct {
	items []Layer
}

// Returns the length of the layer
func (layers *Layers) Length() int {
	return len(layers.items)
}

// Returns the layer level at the specified index
func (layers *Layers) Get(i int) Layer {
	return layers.items[i]
}

// Returns the last layer level
func (layers *Layers) Last() Layer {
	return layers.items[layers.Length() - 1]
}
