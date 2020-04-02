package main

import (
	"iio/pkg/loading"
	"log"
)

func main() {
	bytes, err := loading.LoadMnist()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(*bytes)
	}
}
