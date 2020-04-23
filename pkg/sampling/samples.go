package sampling

import (
	"fmt"
	"math/rand"
)

func NewSamples(length int, maker func(i int) *Sample) *Samples {
	if length < 0 {
		panic(fmt.Sprintf("sampling: length shouldn't be negative, but got %d", length))
	}
	items := make([]*Sample, length)
	for i := 0; i < length; i++ {
		items[i] = maker(i)
	}
	return &Samples{items}
}

type Samples struct {
	items []*Sample
}

func (samples *Samples) Length() int {
	return len(samples.items)
}

func (samples *Samples) To(i int) *Samples {
	if i < 0 || i > samples.Length() {
		panic(fmt.Sprintf("sampling: slice end is out of bounds %d", i))
	}
	items := make([]*Sample, i)
	copy(items, samples.items[:i])
	return &Samples{items}
}

func (samples *Samples) From(i int) *Samples {
	length := samples.Length()
	if i < 0 || i >= length {
		panic(fmt.Sprintf("sampling: slice beginning is out of bounds %d", i))
	}
	items := make([]*Sample, length-i)
	copy(items, samples.items[i:])
	return &Samples{items}
}

func (samples *Samples) Get(i int) *Sample {
	if i < 0 || i >= len(samples.items) {
		panic(fmt.Sprintf("sampling: index is out of bounds %d", i))
	}
	return samples.items[i]
}

func (samples *Samples) Shuffle() *Samples {
	length := samples.Length()
	items := make([]*Sample, length)
	copy(items, samples.items)
	rand.Shuffle(
		length,
		func(i, j int) {
			items[i], items[j] = items[j], items[i]
		},
	)
	return &Samples{items}
}
