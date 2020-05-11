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
	// HTTP request maker, powered by timeout magic.
	client *http.Client

	// A number from (0; 1) indicating relative size of the training
	// images concernedly MNIST overall training set. Used to split
	// this selection into training and validation sets.
	trainingSize float64
}

// Concurrently reads 4 .gz archives, unpacks them, parses and finally
// builds 3 output sample sets. MNIST overall training set should be
// divided to define validation images as well.
func (loader *MNISTLoader) Load() (*sampling.Samples, *sampling.Samples, *sampling.Samples) {
	trainingImageChannel := make(chan []mat.Vector, 1)
	trainingLabelChannel := make(chan []int, 1)
	testImageChannel := make(chan []mat.Vector, 1)
	testLabelChannel := make(chan []int, 1)
	go loader.loadImages("train-images-idx3-ubyte", trainingImageChannel)
	go loader.loadLabels("train-labels-idx1-ubyte", trainingLabelChannel)
	go loader.loadImages("t10k-images-idx3-ubyte", testImageChannel)
	go loader.loadLabels("t10k-labels-idx1-ubyte", testLabelChannel)
	testLabels, trainingLabels := <-testLabelChannel, <-trainingLabelChannel
	testImages, trainingImages := <-testImageChannel, <-trainingImageChannel
	overallSet := loader.makeSamples(trainingImages, trainingLabels)
	trainingIndex := int(loader.trainingSize * float64(overallSet.Length()))
	return overallSet.To(trainingIndex),
		overallSet.From(trainingIndex),
		loader.makeSamples(testImages, testLabels)
}

// Downloads and parses the specified IDX file with the image set
// content. Any error at any stage causes immediate termination.
// Works a set of bytes in IDX format to transform them into
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
func (loader *MNISTLoader) loadImages(filename string, channel chan<- []mat.Vector) {
	pixels, size := loader.loadIDX(filename, 3)
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
	channel <- images
}

// Fetches the target archive and unpacks it straight to the memory.
// Also validates the structure of an IDX byte array. Basically, the
// most significant requirements are appropriate content length,
// content identifier and overall ND-array shape. All the data
// can be taken from a few leading 32-bit integer magic numbers.
// Returns personally ND-array data and the number of items in the
// very first dimension.
func (loader *MNISTLoader) loadIDX(filename string, dimensions int) ([]byte, int) {
	fmt.Printf("Loading %s\n", filename)
	start := time.Now()
	response, err := loader.client.Get(fmt.Sprintf("http://yann.lecun.com/exdb/mnist/%s.gz", filename))
	if err != nil {
		panic(fmt.Sprintf("loading: mnist couldn't get %s\n%v", filename, err))
	}
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		panic(fmt.Sprintf("loading: mnist couldn't decompress %s\n%v", filename, err))
	}
	idx, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(fmt.Sprintf("loading: mnist couldn't read %s\n%v", filename, err))
	}
	err = response.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("loading: mnist couldn't finalize %s\n%v", filename, err))
	}
	idxLength, minLength := len(idx), 4*(dimensions+1)
	if idxLength < minLength {
		panic(fmt.Sprintf("loading: mnist idx should be %d bytes but got %d", minLength, idxLength))
	}
	if idx[0] != 0 || idx[1] != 0 {
		panic(fmt.Sprintf("loading: mnist idx first 2 bytes should be 0 but got %d & %d", idx[0], idx[1]))
	}
	if idx[2] != 8 {
		panic(fmt.Sprintf("loading: mnist idx 3rd byte should be 8 but got %d", idx[2]))
	}
	if idx[3] != byte(dimensions) {
		panic(fmt.Sprintf("loading: mnist idx 4th byte should be %d but got %d", dimensions, idx[3]))
	}
	data, size := idx[minLength:], int(binary.BigEndian.Uint32(idx[4:8]))
	total := size
	for i := 2; i <= dimensions; i++ {
		total *= int(binary.BigEndian.Uint32(idx[i*4 : (i+1)*4]))
	}
	if length := len(data); total != length {
		panic(fmt.Sprintf("loading: mnist idx different lengths %d and %d", total, length))
	}
	fmt.Printf(
		"Loaded %s: %.3f mib (%.3f s)\n",
		filename,
		float64(idxLength)/(1<<20),
		time.Since(start).Seconds(),
	)
	return data, size
}

// Downloads and parses the specified IDX file with the label set.
// Processes an image label IDX slice. Here should be followed
// all the requirements of IDX format plus all the labels should
// be one-digit unsigned integers.
func (loader *MNISTLoader) loadLabels(filename string, channel chan<- []int) {
	bytes, _ := loader.loadIDX(filename, 1)
	labels := make([]int, len(bytes))
	for i, label := range bytes {
		if label > 9 {
			panic(fmt.Sprintf("sampling: mnist idx invalid label %d at %d", label, i))
		}
		labels[i] = int(label)
	}
	channel <- labels
}

// Produces example array - a set of labeled images suitable for
// a network processing.
func (loader *MNISTLoader) makeSamples(images []mat.Vector, labels []int) *sampling.Samples {
	imagesLength, labelsLength := len(images), len(labels)
	if imagesLength != labelsLength {
		panic(fmt.Sprintf("loading: mnist sets have different lengths %d & %d", imagesLength, labelsLength))
	}
	items := make([]*sampling.Sample, imagesLength)
	for i := 0; i < imagesLength; i++ {
		items[i] = sampling.NewSample(images[i], labels[i])
	}
	return sampling.NewSamples(items...)
}
