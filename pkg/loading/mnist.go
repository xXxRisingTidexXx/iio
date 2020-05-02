package loading

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
	"iio/pkg/sampling"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func NewMNISTLoader() *MNISTLoader {
	return &MNISTLoader{&http.Client{Timeout: 10 * time.Second}, 0.9}
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
	client       *http.Client
	trainingSize float64
}

func (loader *MNISTLoader) Load() (*sampling.Samples, *sampling.Samples, *sampling.Samples) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(4)
	trainingImageChannel := make(chan []mat.Vector, 1)
	trainingLabelChannel := make(chan []int, 1)
	testImageChannel := make(chan []mat.Vector, 1)
	testLabelChannel := make(chan []int, 1)
	errChannel := make(chan error, 4)
	go loader.loadImages("train-images-idx3-ubyte", waitGroup, trainingImageChannel, errChannel)
	go loader.loadLabels("train-labels-idx1-ubyte", waitGroup, trainingLabelChannel, errChannel)
	go loader.loadImages("t10k-images-idx3-ubyte", waitGroup, testImageChannel, errChannel)
	go loader.loadLabels("t10k-labels-idx1-ubyte", waitGroup, testLabelChannel, errChannel)
	waitGroup.Wait()
	select {
	case err := <-errChannel:
		return nil, nil, nil, err
	default:
		testLabels, trainingLabels := <-testLabelChannel, <-trainingLabelChannel
		testImages, trainingImages := <-testImageChannel, <-trainingImageChannel
		if err := checkLengths(trainingImages, trainingLabels); err != nil {
			return nil, nil, nil, err
		}
		if err := checkLengths(testImages, testLabels); err != nil {
			return nil, nil, nil, err
		}
		overallSet := makeSamples(trainingImages, trainingLabels)
		trainingIndex := int(loader.trainingSize * float64(overallSet.Length()))
		return overallSet.To(trainingIndex),
			overallSet.From(trainingIndex),
			makeSamples(testImages, testLabels),
			nil
	}
}

// Downloads and parses the specified IDX file with the image set
// content. Any error at any stage causes immediate termination.
func (loader *MNISTLoader) loadImages(
	filename string,
	waitGroup *sync.WaitGroup,
	imageChannel chan<- []mat.Vector,
	errChannel chan<- error,
) {
	defer waitGroup.Done()
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
	fmt.Printf("Loading %s\n", filename)
	start := time.Now()
	response, err := loader.client.Get(fmt.Sprintf("http://yann.lecun.com/exdb/mnist/%s.gz", filename))
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		return nil, err
	}
	idx, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	fmt.Printf(
		"Loaded %s: %.3f mib (%.3f s)\n",
		filename,
		float64(len(idx))/(1<<20),
		time.Since(start).Seconds(),
	)
	return idx, nil
}

// Downloads and parses the specified IDX file with the label set
// content. Any error at any stage causes immediate termination.
func (loader *MNISTLoader) loadLabels(
	filename string,
	waitGroup *sync.WaitGroup,
	labelChannel chan<- []int,
	errChannel chan<- error,
) {
	defer waitGroup.Done()
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
func parseImages(idx []byte) ([]mat.Vector, error) {
	pixels, size, err := checkIDX(idx, 3)
	if err != nil {
		return nil, err
	}
	images, length := make([]mat.Vector, size), len(pixels)/size
	for i := 0; i < size; i++ {
		nonZero := 0
		for j := 0; j < length; j++ {
			if pixels[i*length+j] != 0 {
				nonZero++
			}
		}
		indices, items := make([]int, nonZero), make([]float64, nonZero)
		for j, k := 0, 0; j < length; j++ {
			if pixel := pixels[i*length+j]; pixel != 0 {
				indices[k], items[k] = j, float64(pixel)/255.0
				k++
			}
		}
		images[i] = sparse.NewVector(length, indices, items)
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
func parseLabels(idx []byte) ([]int, error) {
	labels, _, err := checkIDX(idx, 1)
	if err != nil {
		return nil, err
	}
	ints := make([]int, len(labels))
	for i, label := range labels {
		if label > 9 {
			return nil, fmt.Errorf("invalid idx: invalid label %d at index %d", label, i)
		}
		ints[i] = int(label)
	}
	return ints, nil
}

// Checks the lengths of image and label arrays - they must
// be equal to avoid an inconsistency.
func checkLengths(images []mat.Vector, labels []int) error {
	imagesLength, labelsLength := len(images), len(labels)
	if imagesLength != labelsLength {
		return fmt.Errorf("mnist: sets have different lengths %d & %d", imagesLength, labelsLength)
	}
	if imagesLength < 10 {
		return fmt.Errorf("mnist: sets have low length %d", imagesLength)
	}
	return nil
}

// Produces example array - a set of labeled images suitable for
// a network processing.
func makeSamples(images []mat.Vector, labels []int) *sampling.Samples {
	length := len(images)
	items := make([]*sampling.Sample, length)
	for i := 0; i < length; i++ {
		items[i] = sampling.NewSample(images[i], labels[i])
	}
	return sampling.NewSamples(items...)
}
