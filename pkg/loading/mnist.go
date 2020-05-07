package loading

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

func NewMNISTLoaders() (*MNISTLoader, *MNISTLoader) {
	projectPath, err := filepath.Abs(".")
	if err != nil {
		panic(fmt.Errorf("loading: mnist didn't establish project path, %v", err))
	}
	return newMNISTLoader(
			filepath.Join(projectPath, "data", "train-images-idx3-ubyte.idx"),
			filepath.Join(projectPath, "data", "train-labels-idx1-ubyte.idx"),
			60000,
		),
		newMNISTLoader(
			filepath.Join(projectPath, "data", "t10k-images-idx3-ubyte.idx"),
			filepath.Join(projectPath, "data", "t10k-labels-idx1-ubyte.idx"),
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
		panic(fmt.Errorf("loading: mnist got too low batch size %d", size))
	}
	if loader.position >= loader.length {
		panic(fmt.Errorf("loading: mnist batching ended"))
	}
	if difference := loader.length - loader.position; difference < size {
		size = difference
	}

	batch := make([]*Sample, size)
	loader.position += size
	return batch
}

func (loader *MNISTLoader) readIDX(size int) ([][]byte, [][]byte) {
	indices := loader.indices[loader.position : loader.position+size]
	images, labels := make([][]byte, size), make([][]byte, size)
	imagesFile, err := os.Open(loader.imageFilePath)
	if err != nil {
		panic(fmt.Errorf("loading: mnist images didn't open, %v", err))
	}
	labelsFile, err := os.Open(loader.labelFilePath)
	if err != nil {
		panic(fmt.Errorf("loading: mnist labels didn't open, %v", err))
	}
	for i := 0; i < size; i++ {
		images[i], labels[i] = make([]byte, loader.imageExtent), make([]byte, loader.labelExtent)
		read, err := imagesFile.ReadAt(images[i], int64(loader.imageShift+loader.imageExtent*indices[i]))
		if err != nil {
			panic(fmt.Errorf("loading: mnist image reading failed, %v", err))
		}
		if read != loader.imageExtent {
			panic(fmt.Errorf("loading: mnist image extent mismatch, %d", read))
		}
		read, err = labelsFile.ReadAt(labels[i], int64(loader.labelShift+loader.labelExtent*indices[i]))
		if err != nil {
			panic(fmt.Errorf("loading: mnist label reading failed, %v", err))
		}
		if read != loader.labelExtent {
			panic(fmt.Errorf("loading: mnist label extent mismatch, %d", read))
		}
	}
	if err := imagesFile.Close(); err != nil {
		panic(fmt.Errorf("loading: mnist images didn't close, %v", err))
	}
	if err := labelsFile.Close(); err != nil {
		panic(fmt.Errorf("loading: mnist labels didn't close, %v", err))
	}
	return images, labels
}
