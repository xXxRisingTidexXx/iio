package sampling

import (
	"fmt"
	"math/rand"
)

func NewSamples(length int, maker func(int) *Sample) *Samples {
	if length < 0 {
		panic(fmt.Sprintf("sampling: length shouldn't be negative, but got %d", length))
	}
	if maker == nil {
		panic(fmt.Sprintf("sampling: maker can't be nil"))
	}
	items := make([]*Sample, length)
	for i := 0; i < length; i++ {
		items[i] = maker(i)
		if items[i] == nil {
			panic(fmt.Sprintf("sampling: %d sample is nil", i))
		}
	}
	return &Samples{items, length, 0}
}

type Samples struct {
	items    []*Sample
	length   int
	position int
}

func (samples *Samples) Length() int {
	return samples.length
}

func (samples *Samples) To(i int) *Samples {
	if i < 0 || i > samples.length {
		panic(fmt.Sprintf("sampling: slice end is out of bounds %d", i))
	}
	items := make([]*Sample, i)
	if i > 0 {
		copy(items, samples.items[:i])
	}
	return &Samples{items, i, 0}
}

func (samples *Samples) From(i int) *Samples {
	if i < 0 || i > samples.length {
		panic(fmt.Sprintf("sampling: slice beginning is out of bounds %d", i))
	}
	newLength := samples.length - i
	items := make([]*Sample, newLength)
	if newLength > 0 {
		copy(items, samples.items[i:])
	}
	return &Samples{items, newLength, 0}
}

func (samples *Samples) Get(i int) *Sample {
	if i < 0 || i >= samples.length {
		panic(fmt.Sprintf("sampling: index is out of bounds %d", i))
	}
	return samples.items[i]
}

func (samples *Samples) Shuffle() *Samples {
	items := make([]*Sample, samples.length)
	copy(items, samples.items)
	rand.Shuffle(
		samples.length,
		func(i, j int) {
			items[i], items[j] = items[j], items[i]
		},
	)
	return &Samples{items, samples.length, 0}
}

func (samples *Samples) Next() bool {
	isAvailable := samples.position < samples.length
	if !isAvailable {
		samples.position = 0
	}
	return isAvailable
}

func (samples *Samples) Batch(size int) *Samples {
	if size < 1 {
		panic(fmt.Sprintf("sampling: too low batch size %d", size))
	}
	if !samples.Next() {
		panic(fmt.Sprintf("sampling: iteration ended"))
	}
	offset := size
	if difference := samples.length - samples.position; difference < size {
		offset = difference
	}
	batch := &Samples{samples.items[samples.position : samples.position+offset], offset, 0}
	samples.position += offset
	return batch
}
