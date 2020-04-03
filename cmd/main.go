package main

import (
	"iio/pkg/loading"
	"log"
)

func main() {
	bytes, err := loading.LoadMNIST()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(*bytes)
	}
}
