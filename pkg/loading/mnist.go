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
		panic(fmt.Errorf("loading: project path didn't established, %v", err))
	}
	return newMNISTLoader(
			filepath.Join(projectPath, "data", "train-images-idx3-ubyte.idx"),
			filepath.Join(projectPath, "data", "train-labels-idx1-ubyte.idx"),
		),
		newMNISTLoader(
			filepath.Join(projectPath, "data", "t10k-images-idx3-ubyte.idx"),
			filepath.Join(projectPath, "data", "t10k-labels-idx1-ubyte.idx"),
		)
}

func newMNISTLoader(imagesFilePath, labelsFilePath string) *MNISTLoader {

	return &MNISTLoader{}
}

type MNISTLoader struct {
	imagesFilePath string
	labelsFilePath string
	length         int
	size           int
	indices        []int
	position       int
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
		panic(fmt.Errorf("loading: too low batch size %d", size))
	}
	if loader.position >= loader.length {
		panic(fmt.Errorf("loading: batching ended"))
	}
	imagesFile, err := os.Open(loader.imagesFilePath)
	if err != nil {
		panic(fmt.Errorf("loading: images didn't open, %v", err))
	}
	labelsFile, err := os.Open(loader.labelsFilePath)
	if err != nil {
		panic(fmt.Errorf("loading: labels didn't open, %v", err))
	}


	offset := size
	if difference := loader.length - loader.position; difference < size {
		offset = difference
	}
	batch := make([]*Sample, offset)

	if err := imagesFile.Close(); err != nil {
		panic(fmt.Errorf("loading: images didn't close, %v", err))
	}
	if err := labelsFile.Close(); err != nil {
		panic(fmt.Errorf("loading: labels didn't close, %v", err))
	}
	loader.position += offset
	return batch
}
