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
	"time"
)

func NewMNISTLoader() Loader {
	return &MNISTLoader{&http.Client{Timeout: 5 * time.Second}, &sync.WaitGroup{}}
}

// Downloads the whole MNIST database images with a single batch.
// The homepage of the DB is http://yann.lecun.com/exdb/mnist/ .
// 4 concurrent requests should fetch .gz archives, decompress them
// and convert into numerical data structures. Actually, on the
// 14 April 2020, sizes of the image and label files of the training
// and test sets should equal 44.9 mib, 57 kib, 7.5 mib and 10 kib
// respectively. All files are represented by IDX format which is
// very suitable for ND-array transfer.
type MNISTLoader struct {
	client    *http.Client
	waitGroup *sync.WaitGroup
}

func (loader *MNISTLoader) Load() ([]*Example, []*Example, error) {
	trainingImageChannel := make(chan []vectors.Vector, 1)
	trainingLabelChannel := make(chan []byte, 1)
	testImageChannel := make(chan []vectors.Vector, 1)
	testLabelChannel := make(chan []byte, 1)
	errChannel := make(chan error, 4)
	loader.waitGroup.Add(4)
	go loader.loadImages("train-images-idx3-ubyte", trainingImageChannel, errChannel)
	go loader.loadLabels("train-labels-idx1-ubyte", trainingLabelChannel, errChannel)
	go loader.loadImages("t10k-images-idx3-ubyte", testImageChannel, errChannel)
	go loader.loadLabels("t10k-labels-idx1-ubyte", testLabelChannel, errChannel)
	loader.waitGroup.Wait()
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

// Downloads and parses the specified IDX file with the image set
// content. Any error at any stage causes immediate termination.
func (loader *MNISTLoader) loadImages(
	filename string,
	imageChannel chan<- []vectors.Vector,
	errChannel chan<- error,
) {
	defer loader.waitGroup.Done()
	idx, err := loader.getAndDecompressIDX(filename)
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

// Fetches the target archive and unpacks it straight to the memory.
func (loader *MNISTLoader) getAndDecompressIDX(filename string) ([]byte, error) {
	log.Printf("Loading %s\n", filename)
	response, err := loader.client.Get(fmt.Sprintf("http://yann.lecun.com/exdb/mnist/%s.gz", filename))
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

// Downloads and parses the specified IDX file with the label set
// content. Any error at any stage causes immediate termination.
func (loader *MNISTLoader) loadLabels(
	filename string,
	labelChannel chan<- []byte,
	errChannel chan<- error,
) {
	defer loader.waitGroup.Done()
	idx, err := loader.getAndDecompressIDX(filename)
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

// Accepts a set of bytes in IDX format to transform them into
// more "mathematical" data structure - vector.Vector . The
// specification of MNIST declares images in the form of the 3D
// tensor (60000 images x 28 pixels width x 28 pixels height),
// where each item is an int [0; 255]; the higher num, the lighter
// pixel. Their general amount yields the contour of the regular
// cipher. Here each image is converted from the "2D view" into
// 1D array where each element is in range [0; 1] - all pictures
// are flattened and divided by 255 to obtain activation vector
// for further computations. All images are read in C-style - i.e.
// row-by-row or row-wise.
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

// Validates the structure of an IDX byte array. Basically, the
// most significant requirements are appropriate content length,
// content identifier and overall ND-array shape. All the data
// can be taken from a few leading 32-bit integer magic numbers.
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

// Processes an image label IDX slice. Here should be followed
// all the requirements of IDX format plus all the labels should
// be one-digit unsigned integers.
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

// Checks the lengths of image and label arrays - they must
// be equal to avoid an inconsistency.
func compareLengths(images []vectors.Vector, labels []byte) error {
	imagesLength, labelsLength := len(images), len(labels)
	if imagesLength == labelsLength {
		return nil
	}
	return fmt.Errorf("mnist: sets have different lengths %d & %d", imagesLength, labelsLength)
}

// Produces example array - a set of labeled images suitable for
// a network processing.
func makeExamples(images []vectors.Vector, labels []byte) []*Example {
	length := len(labels)
	examples := make([]*Example, length)
	for i := 0; i < length; i++ {
		examples[i] = &Example{images[i], labels[i]}
	}
	return examples
}
