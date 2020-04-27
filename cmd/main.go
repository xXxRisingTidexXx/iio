package main

import (
	"fmt"
	"iio/pkg/loading"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	loader := loading.NewMNISTLoader()
	if trainingSamples, validationSamples, testSamples, err := loader.Load(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Training set length: %d\n", trainingSamples.Length())
		fmt.Printf("Validation set length: %d\n", validationSamples.Length())
		fmt.Printf("Test set length: %d\n", testSamples.Length())
		fmt.Println(trainingSamples.Get(12047).Activations)
		fmt.Println(validationSamples.Get(3021).Activations)
		fmt.Println(testSamples.Get(9541).Activations)
	}
}
