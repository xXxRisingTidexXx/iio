package main

import (
	"iio/pkg/loading"
	"log"
)

func main() {
	trainingExamples, testExamples, err := loading.LoadMNIST()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(len(trainingExamples))
		log.Println(trainingExamples[0].Image)
		log.Println(trainingExamples[0].Label)
		log.Println(len(testExamples))
		log.Println(testExamples[0].Image)
		log.Println(testExamples[0].Label)
	}
}
