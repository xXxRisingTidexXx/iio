package loading

type Loader interface {
	Length() int
	Shuffle()
	Next() bool
	Batch(int) []*Sample
}
