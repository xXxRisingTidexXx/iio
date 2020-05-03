package sampling

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math/rand"
)

func NewSamples(items ...*Sample) *Samples {
	if items == nil {
		items = make([]*Sample, 0)
	}
	for i, item := range items {
		if item == nil {
			panic(fmt.Sprintf("sampling: sample at %d is nil", i))
		}
		if i > 1 && item.activations.Len() != items[i-1].activations.Len() {
			panic(fmt.Sprintf("sampling: sample at %d has uncommon length", i))
		}
	}
	return &Samples{items, len(items), 0}
}

type Samples struct {
	items    []*Sample
	length   int
	position int
}

func (samples *Samples) Equal(other *Samples) bool {
	return samples == other ||
		other != nil &&
			cmp.Equal(samples.items, other.items) &&
			samples.length == other.length &&
			samples.position == other.position
}

func (samples *Samples) String() string {
	return fmt.Sprintf("{%v %d %d}", samples.items, samples.length, samples.position)
}

func (samples *Samples) Length() int {
	return samples.length
}

func (samples *Samples) To(i int) *Samples {
	end := i
	if end < 0 {
		end = 0
	} else if end > samples.length {
		end = samples.length
	}
	items := make([]*Sample, end)
	if end > 0 {
		copy(items, samples.items[:end])
	}
	return &Samples{items, end, 0}
}

func (samples *Samples) From(i int) *Samples {
	start := i
	if start < 0 {
		start = 0
	} else if start > samples.length {
		start = samples.length
	}
	length := samples.length - start
	items := make([]*Sample, length)
	if length > 0 {
		copy(items, samples.items[start:])
	}
	return &Samples{items, length, 0}
}

func (samples *Samples) Get(i int) *Sample {
	if i < 0 || i >= samples.length {
		panic(fmt.Sprintf("sampling: index %d is out of bounds [%d; %d)", i, 0, samples.length))
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
		panic(fmt.Sprintf("sampling: batching ended"))
	}
	offset := size
	if difference := samples.length - samples.position; difference < size {
		offset = difference
	}
	batch := &Samples{samples.items[samples.position : samples.position+offset], offset, 0}
	samples.position += offset
	return batch
}
