package loading

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"iio/pkg/vectors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func LoadMNIST() ([]*Example, []*Example, error) {
	waitGroup := sync.WaitGroup{}
	trainingImageChannel := make(chan []vectors.Vector, 1)
	trainingLabelChannel := make(chan []byte, 1)
	testImageChannel := make(chan []vectors.Vector, 1)
	testLabelChannel := make(chan []byte, 1)
	errChannel := make(chan error, 4)
	waitGroup.Add(4)
	go loadImages("train-images-idx3-ubyte", &waitGroup, trainingImageChannel, errChannel)
	go loadLabels("train-labels-idx1-ubyte", &waitGroup, trainingLabelChannel, errChannel)
	go loadImages("t10k-images-idx3-ubyte", &waitGroup, testImageChannel, errChannel)
	go loadLabels("t10k-labels-idx1-ubyte", &waitGroup, testLabelChannel, errChannel)
	waitGroup.Wait()
	select {
	case err := <-errChannel:
		return nil, nil, err
	default:
		trainingImages, trainingLabels := <-trainingImageChannel, <-trainingLabelChannel
		testImages, testLabels := <-testImageChannel, <-testLabelChannel
		if err := compareLengths(trainingImages, trainingLabels); err != nil {
			return nil, nil, err
		}
		if err := compareLengths(testImages, testLabels); err != nil {
			return nil, nil, err
		}
		return makeExamples(trainingImages, trainingLabels), makeExamples(testImages, testLabels), nil
	}
}

func loadImages(
	filename string,
	waitGroup *sync.WaitGroup,
	imageChannel chan<- []vectors.Vector,
	errChannel chan<- error,
) {
	defer waitGroup.Done()
	idx, err := getAndDecompressIDX(filename)
	if err != nil {
		errChannel <- fmt.Errorf("mnist: %s: %v", filename, err)
		return
	}
	images, err := parseImages(idx)
	if err != nil {
		errChannel <- fmt.Errorf("mnist: %s: %v", filename, err)
		return
	}
	imageChannel <- images
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
	log.Printf("Downloaded %s: %.3f mib\n", filename, float64(len(idx))/(1<<20))
	return idx, nil
}

func parseImages(idx []byte) ([]vectors.Vector, error) {
	pixels, size, err := checkIDX(idx, 3)
	if err != nil {
		return nil, err
	}
	images, length := make([]vectors.Vector, size), len(pixels)/size
	for i := 0; i < size; i += length {
		items := make([]float64, length)
		for j := 0; j < length; j++ {
			items[j] = float64(pixels[i+j]) / 255.0
		}
		images[i] = vectors.Vectorize(items)
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
		return nil, 0, fmt.Errorf("invalid idx: 4th byte should be %d but got %d", dimensions, idx[3])
	}
	data, size := idx[minLength:], int(binary.BigEndian.Uint32(idx[4:8]))
	total := size
	for i := 2; i <= dimensions; i++ {
		total *= int(binary.BigEndian.Uint32(idx[i*4 : (i+1)*4]))
	}
	if length := len(data); total != length {
		return nil, 0, fmt.Errorf("invalid idx: different lengths %d and %d", total, length)
	}
	return data, size, nil
}

func loadLabels(
	filename string,
	waitGroup *sync.WaitGroup,
	labelChannel chan<- []byte,
	errChannel chan<- error,
) {
	defer waitGroup.Done()
	idx, err := getAndDecompressIDX(filename)
	if err != nil {
		errChannel <- fmt.Errorf("mnist: %s: %v", filename, err)
		return
	}
	labels, err := parseLabels(idx)
	if err != nil {
		errChannel <- fmt.Errorf("mnist: %s: %v", filename, err)
		return
	}
	labelChannel <- labels
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

func compareLengths(images []vectors.Vector, labels []byte) error {
	imagesLength, labelsLength := len(images), len(labels)
	if imagesLength == labelsLength {
		return nil
	}
	return fmt.Errorf("mnist: sets have different lengths %d & %d", imagesLength, labelsLength)
}

func makeExamples(images []vectors.Vector, labels []byte) []*Example {
	length := len(labels)
	examples := make([]*Example, length)
	for i := 0; i < length; i++ {
		examples[i] = &Example{images[i], labels[i]}
	}
	return examples
}
