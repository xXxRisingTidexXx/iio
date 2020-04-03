package loading

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func LoadMNIST() (*[]float64, error) {
	content, _ := getAndDecompressIDX("t10k-labels-idx1-ubyte")
	labels, _ := parseLabels(content)
	log.Info(*labels)
	log.Info(len(*labels))
	return &[]float64{}, nil
}

func getAndDecompressIDX(filename string) (*[]byte, error) {
	log.Infof("Loading %s", filename)
	response, err := http.Get(fmt.Sprintf("http://yann.lecun.com/exdb/mnist/%s.gz", filename))
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		return nil, err
	}
	log.Infof("Decompressing %s", filename)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Infof("Loading of %s completed: %d\n", filename, len(content))
	return &content, nil
}

func parseImages(content *[]byte) (*[]*[]float64, error) {
	return nil, nil
}

func parseLabels(content *[]byte) (*[]byte, error) {
	idx := *content
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
	return &labels, nil
}
