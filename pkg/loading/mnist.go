package loading

import (
	"fmt"
	"github.com/james-bowman/sparse"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
)

func NewMNISTLoaders() (*MNISTLoader, *MNISTLoader) {
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("loading: root dir wasn't instantiated")
	}
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(filePath)))
	return newMNISTLoader(
			filepath.Join(rootDir, "data", "train-images-idx3-ubyte.idx"),
			filepath.Join(rootDir, "data", "train-labels-idx1-ubyte.idx"),
			60000,
		),
		newMNISTLoader(
			filepath.Join(rootDir, "data", "t10k-images-idx3-ubyte.idx"),
			filepath.Join(rootDir, "data", "t10k-labels-idx1-ubyte.idx"),
			10000,
		)
}

func newMNISTLoader(imageFilePath, labelFilePath string, length int) *MNISTLoader {
	indices := make([]int, length)
	for i := 0; i < length; i++ {
		indices[i] = i
	}
	return &MNISTLoader{imageFilePath, 16, 28 * 28, labelFilePath, 8, 1, length, indices, 0}
}

type MNISTLoader struct {
	imageFilePath string
	imageShift    int
	imageExtent   int
	labelFilePath string
	labelShift    int
	labelExtent   int
	length        int
	indices       []int
	position      int
}

func (loader *MNISTLoader) Length() int {
	return loader.length
}

func (loader *MNISTLoader) Shuffle() {
	rand.Shuffle(
		loader.length,
		func(i, j int) {
			loader.indices[i], loader.indices[j] = loader.indices[j], loader.indices[i]
		},
	)
}

func (loader *MNISTLoader) Next() bool {
	isAvailable := loader.position < loader.length
	if !isAvailable {
		loader.position = 0
	}
	return isAvailable
}

func (loader *MNISTLoader) Batch(size int) []*Sample {
	if size < 1 {
		panic(fmt.Sprintf("loading: mnist got too low batch size %d", size))
	}
	if loader.position >= loader.length {
		panic("loading: mnist batching ended")
	}
	if difference := loader.length - loader.position; difference < size {
		size = difference
	}
	images, labels := loader.readIDX(size)
	batch := make([]*Sample, size)
	for i := 0; i < size; i++ {
		nonZeroCount := 0
		for j := 0; j < loader.imageExtent; j++ {
			if images[i][j] != 0 {
				nonZeroCount++
			}
		}
		indices, data := make([]int, nonZeroCount), make([]float64, nonZeroCount)
		for j, k := 0, 0; j < loader.imageExtent; j++ {
			if images[i][j] != 0 {
				indices[k], data[k] = j, float64(images[i][j])/255
				k++
			}
		}
		batch[i] = NewSample(sparse.NewVector(loader.imageExtent, indices, data), int(labels[i][0]))
	}
	loader.position += size
	return batch
}

func (loader *MNISTLoader) readIDX(size int) ([][]byte, [][]byte) {
	indices := loader.indices[loader.position : loader.position+size]
	images, labels := make([][]byte, size), make([][]byte, size)
	imagesFile, err := os.Open(loader.imageFilePath)
	if err != nil {
		panic(fmt.Sprintf("loading: mnist images didn't open, %v", err))
	}
	labelsFile, err := os.Open(loader.labelFilePath)
	if err != nil {
		panic(fmt.Sprintf("loading: mnist labels didn't open, %v", err))
	}
	for i := 0; i < size; i++ {
		images[i], labels[i] = make([]byte, loader.imageExtent), make([]byte, loader.labelExtent)
		read, err := imagesFile.ReadAt(images[i], int64(loader.imageShift+loader.imageExtent*indices[i]))
		if err != nil {
			panic(fmt.Sprintf("loading: mnist image reading failed, %v", err))
		}
		if read != loader.imageExtent {
			panic(fmt.Sprintf("loading: mnist image extent mismatch, %d", read))
		}
		read, err = labelsFile.ReadAt(labels[i], int64(loader.labelShift+loader.labelExtent*indices[i]))
		if err != nil {
			panic(fmt.Sprintf("loading: mnist label reading failed, %v", err))
		}
		if read != loader.labelExtent {
			panic(fmt.Sprintf("loading: mnist label extent mismatch, %d", read))
		}
	}
	if err := imagesFile.Close(); err != nil {
		panic(fmt.Sprintf("loading: mnist images didn't close, %v", err))
	}
	if err := labelsFile.Close(); err != nil {
		panic(fmt.Sprintf("loading: mnist labels didn't close, %v", err))
	}
	return images, labels
}
