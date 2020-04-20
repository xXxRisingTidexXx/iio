package main

import (
	"fmt"
	"iio/pkg/loading"
)

func main() {
	if trainingSamples, validationSamples, testSamples, err := loading.NewMNISTLoader().Load(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Training set length: %d\n", len(trainingSamples))
		fmt.Printf("Validation set length: %d\n", len(validationSamples))
		fmt.Printf("Test set length: %d\n", len(testSamples))
	}
}
