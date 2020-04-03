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
	idx, _ := getAndDecompressIDX("t10k-labels-idx1-ubyte")
	labels, _ := parseLabels(idx)
	log.Println(labels)
	log.Println(len(labels))
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
	if len(idx) < 16 {
		return nil, fmt.Errorf("invalid idx: should contain at least 16 bytes but got %d", len(idx))
	}
	if idx[0] != 0 || idx[1] != 0 {
		return nil, fmt.Errorf("invalid idx: first 2 bytes should be 0 but got %d & %d", idx[0], idx[1])
	}
	if idx[2] != 8 {
		return nil, fmt.Errorf("invalid idx: 3rd byte should be 8 but got %d", idx[2])
	}
	if idx[3] != 3 {
		return nil, fmt.Errorf("invalid idx: 4th byte should be 3 but got %d", idx[2])
	}
	pixels := idx[16:]
	size := binary.BigEndian.Uint64(idx[4:8])
	rows := binary.BigEndian.Uint64(idx[8:12])
	columns := binary.BigEndian.Uint64(idx[8:12])
	if size*rows*columns != uint64(len(pixels)) {
		return nil,
			fmt.Errorf(
				"invalid idx: different lengths %d x %d x %d and %d", size, rows, columns, len(pixels),
			)
	}
	images := make([]vectors.Vector, size)
	for i := uint64(0); i < size; i++ {
		items := make([]float64, rows*columns)
		for j := uint64(0); j < rows; j++ {
			for k := uint64(0); k < columns; k++ {
				items[j * rows + k] = float64(pixels[i * size + j * rows + k]) / 255.0
			}
		}
		images[i] = vectors.Vectorize(items)
	}
	return images, nil
}

func parseLabels(idx []byte) ([]byte, error) {
	if len(idx) < 8 {
		return nil, fmt.Errorf("invalid idx: should contain at least 8 bytes but got %d", len(idx))
	}
	if idx[0] != 0 || idx[1] != 0 {
		return nil, fmt.Errorf("invalid idx: first 2 bytes should be 0 but got %d & %d", idx[0], idx[1])
	}
	if idx[2] != 8 {
		return nil, fmt.Errorf("invalid idx: 3rd byte should be 8 but got %d", idx[2])
	}
	if idx[3] != 1 {
		return nil, fmt.Errorf("invalid idx: 4th byte should be 1 but got %d", idx[2])
	}
	labels := idx[8:]
	if size := binary.BigEndian.Uint32(idx[4:8]); size != uint32(len(labels)) {
		return nil, fmt.Errorf("invalid idx: different lengths %d and %d", size, len(labels))
	}
	for i, label := range labels {
		if label > 9 {
			return nil, fmt.Errorf("invalid idx: invalid label %d at index %d", label, i)
		}
	}
	return labels, nil
}
