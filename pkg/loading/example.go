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
		return fmt.Sprintf("%v , (%d,) - %d", image, length, example.Label)
	}
	return fmt.Sprintf(
		"[%.4f %.4f %.4f ... %.4f %.4f %.4f] - (%d), %d",
		image[0],
		image[1],
		image[2],
		image[length-3],
		image[length-2],
		image[length-1],
		length,
		example.Label,
	)
}
