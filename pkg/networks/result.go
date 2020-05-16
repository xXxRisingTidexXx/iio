package networks

import (
	"fmt"
)

func newResult(actual, ideal int) *result {
	if actual < 0 {
		panic(fmt.Sprintf("networks: result got invalid actual, %d", actual))
	}
	if ideal < 0 {
		panic(fmt.Sprintf("networks: result got invalid ideal, %d", ideal))
	}
	return &result{actual, ideal}
}

type result struct {
	actual int
	ideal  int
}
