package sampling

import "math/rand"

type Samples struct {
	items []*Sample
}

func (samples *Samples) Length() int {
	return len(samples.items)
}

func (samples *Samples) Get(i int) *Sample {
	return samples.items[i]
}

func (samples *Samples) Shuffle() *Samples {
	length := len(samples.items)
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
