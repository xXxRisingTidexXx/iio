package loading

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

func NewMNISTLoader() *MNISTLoader {
	return &MNISTLoader{}
}

type MNISTLoader struct {
	client http.Client
}

func LoadMnist() (*[]float64, error) {
	log.Println("Loading training images...")
	response, err := http.Get("http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz")
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		return nil, err
	}
	log.Println("Decompressing training images...")
	images, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Printf("Loading successfully completed: %d\n", len(images))
	return &[]float64{}, nil
}

func loadAndDecompress(filename string) (*[]byte, error) {

}
