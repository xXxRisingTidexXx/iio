package layers

import (
	"fmt"
)

func NewSchema(kind Kind, size int) *Schema {
	if size < 1 {
		panic(fmt.Sprintf("layers: schema got invalid size, %d", size))
	}
	return &Schema{kind, size}
}

type Schema struct {
	Kind Kind
	Size int
}
