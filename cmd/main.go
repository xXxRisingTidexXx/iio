package main

import (
	"fmt"
	"iio/pkg/loading"
)

func main() {
	_, testLoader := loading.NewMNISTLoaders()
	sample := testLoader.Batch(10)[9]
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			x := sample.Data().AtVec(i*28 + j)
			if x > 0.0 {
				fmt.Print(fmt.Sprintf("%.2f ", x))
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
	fmt.Println(sample.Label())
}
