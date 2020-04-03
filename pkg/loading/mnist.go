package loading

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"iio/pkg/vectors"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadMNIST() (*[]float64, error) {
	idx, _ := getAndDecompressIDX("t10k-images-idx3-ubyte")
	images, _ := parseImages(idx)
	idx, _ = getAndDecompressIDX("t10k-labels-idx1-ubyte")
	labels, _ := parseLabels(idx)
	log.Println(len(images))
	log.Println(images[0])
	log.Println(len(labels))
	log.Println(labels[0])
	return &[]float64{}, nil
}

func getAndDecompressIDX(filename string) ([]byte, error) {
	log.Printf("Loading %s\n", filename)
	response, err := http.Get(fmt.Sprintf("http://yann.lecun.com/exdb/mnist/%s.gz", filename))
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Decompressing %s\n", filename)
	idx, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Printf("Loading of %s completed: %d\n", filename, len(idx))
	return idx, nil
}

func parseImages(idx []byte) ([]vectors.Vector, error) {
	pixels, size, err := checkIDX(idx, 3)
	if err != nil {
		return nil, err
	}
	images, length := make([]vectors.Vector, size), len(pixels) / size
	for i := 0; i < size; i += length {
		activations := make([]float64, length)
		for j := 0; j < length; j++ {
			activations[j] = float64(pixels[i + j]) / 255.0
		}
		images[i] = vectors.Vectorize(activations)
	}
	return images, nil
}

func checkIDX(idx []byte, dimensions int) ([]byte, int, error) {
	minLength := 4 * (dimensions + 1)
	if len(idx) < minLength {
		return nil, 0, fmt.Errorf("invalid idx: too short - %d bytes, expected %d", len(idx), minLength)
	}
	if idx[0] != 0 || idx[1] != 0 {
		return nil, 0, fmt.Errorf("invalid idx: first 2 bytes should be 0 but got %d & %d", idx[0], idx[1])
	}
	if idx[2] != 8 {
		return nil, 0, fmt.Errorf("invalid idx: 3rd byte should be 8 but got %d", idx[2])
	}
	if idx[3] != byte(dimensions) {
		return nil, 0, fmt.Errorf("invalid idx: 4th byte should be %d but got %d", dimensions, idx[2])
	}
	data, size := idx[minLength:], int(binary.BigEndian.Uint32(idx[4:8]))
	total := size
	for i := 2; i <= dimensions; i++ {
		total *= int(binary.BigEndian.Uint32(idx[i * 4: (i + 1) * 4]))
	}
	if length := len(data); total != length {
		return nil, 0, fmt.Errorf("invalid idx: different lengths %d and %d", total, length)
	}
	return data, size, nil
}

func parseLabels(idx []byte) ([]byte, error) {
	labels, _, err := checkIDX(idx, 1)
	if err != nil {
		return nil, err
	}
	for i, label := range labels {
		if label > 9 {
			return nil, fmt.Errorf("invalid idx: invalid label %d at index %d", label, i)
		}
	}
	return labels, nil
}
