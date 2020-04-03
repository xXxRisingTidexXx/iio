package loading

import "fmt"

type Example struct {
	Image *[]float64
	Label byte
}

func (example *Example) String() string {
	image := *example.Image
	length := len(image)
	if length <= 10 {
		return fmt.Sprintf("%v - (%d), %d", image, length, example.Label)
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
