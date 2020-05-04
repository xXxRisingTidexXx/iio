package sampling

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math/rand"
)

// Constructs a sample sequence from the provided variadic elements.
// Replaces nil slices with zero-length slices. All examples must be
// non-nil structs with the same length activation vectors.
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

// Sample collection wrapper. Supplies convenient interface for
// common sequence operations like index access, slicing and length.
// Also provides specific tools for shuffling (random element mixing)
// and iteration. Samples was thought to be an immutable collection,
// but iteration need made adjustments, so now the sequence has a
// mutable state. But even though samples can be easily used in an
// exception-safe manner, whose rules are described below.
type Samples struct {
	// Learning object pointer slice.
	items []*Sample

	// The sequence length shortcut field used to avoid permanent
	// `len(samples.items)` calls.
	length int

	// Stuff like a DB cursor object pointing to the beginning of
	// the available for iteration elements.
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

// Public collection size accessor.
func (samples *Samples) Length() int {
	return samples.length
}

// Truncates the underlying array up to the specified index
// exclusively. If the provided position is too high, the whole
// collection will be copied. If it's too low, empty samples will be
// returned.
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

// Truncates the underlying array from the specified position
// inclusively. If the provided position is too high, empty samples
// will be returned. If it's too low, the whole collection will be
// copied.
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

// Returns the element at the specified position. In case of the
// boundary violation will produce a panic. The struct tends to
// be almost immutable, so the returned sample shouldn't be changed
// in any case.
func (samples *Samples) Get(i int) *Sample {
	if i < 0 || i >= samples.length {
		panic(fmt.Sprintf("sampling: index %d is out of bounds [%d; %d)", i, 0, samples.length))
	}
	return samples.items[i]
}

// Randomly rearranges the underlying collection producing a new
// selection of the same size. To apply a truly pseudorandom
// generation, https://golang.org/pkg/math/rand/#Seed command should
// be ordered at the application top level.
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

// One of two important callables needed for samples' iteration
// facade. `Cause go doesn't provide a built-in way to range over
// custom collections, local "patches" may be applied to solve
// concrete design problems. In our case sample array should provide
// a convenient way to range over itself with a defined subsequence
// size. Therefore client would have an ability to split the incoming
// set into sub-arrays whose personal length would be less than or
// equal the pre-defined number. It's a useful tool for mini-batch
// gradient descent, where samples are widely used.
//
// So, specifically this method should be used as a condition at the
// top of a for loop ranging over the collection with a certain batch
// size. After an iteration end resets the position index to 0 to
// supply an opportunity for a one more iteration.
func (samples *Samples) Next() bool {
	isAvailable := samples.position < samples.length
	if !isAvailable {
		samples.position = 0
	}
	return isAvailable
}

// Ordinary subsequence extraction method, smth like `cursor.fetch()`
// function in DB cursor implementations. Fetches sub-samples from
// the original ones, whose length doesn't exceed the specified batch
// size. The only safe way to use this method is:
//   samples := sampling.NewSamples(...)
//   for samples.Next() {
//       batch := samples.Batch(10)
//       ...
//   }
// The bunch consisting of `Sample.Next()` and `Sample.Batch()`
// methods is the only way guaranteeing the safe iteration. Don't
// use the batching function out of a loop wrapped by `Sample.Next()`
// instruction - it carries out position check and reset. Also it
// automatically protects a client from the empty sample batching
// (which is forbidden). Avoid repeatable `Sample.Batch()` calls,
// `cause it may break the internal iteration mechanism. "Repeatable"
// means that the batching function must be called no more than
// once per iteration.
func (samples *Samples) Batch(size int) *Samples {
	if size < 1 {
		panic(fmt.Sprintf("sampling: too low batch size %d", size))
	}
	if samples.position >= samples.length {
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
