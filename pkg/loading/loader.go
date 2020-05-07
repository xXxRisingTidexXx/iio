package loading

type Loader interface {
	Shuffle()
	Next() bool
	Batch(int) []*Sample
}
