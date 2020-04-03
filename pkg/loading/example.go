package loading

import (
	"fmt"
	"iio/pkg/vectors"
)

type Example struct {
	Image vectors.Vector
	Label byte
}

func (example *Example) String() string {
	length := example.Image.Length()
	if length <= 10 {
		return fmt.Sprintf("%v , (%d,) - %d", example.Image, length, example.Label)
	}
	return fmt.Sprintf(
		"[%.4f %.4f %.4f ... %.4f %.4f %.4f] - (%d), %d",
		example.Image.Get(0),
		example.Image.Get(1),
		example.Image.Get(2),
		example.Image.Get(length-3),
		example.Image.Get(length-2),
		example.Image.Get(length-1),
		length,
		example.Label,
	)
}
