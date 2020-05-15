package estimate

import (
	"fmt"
)

type Record struct {
	Precision float64
	Recall    float64
	F1Score   float64
	Support   int
}

func (record *Record) Equal(other *Record) bool {
	return other != nil &&
		record.Precision == other.Precision &&
		record.Recall == other.Recall &&
		record.F1Score == other.F1Score &&
		record.Support == other.Support
}

func (record *Record) String() string {
	return fmt.Sprintf("{%f %f %f %d}", record.Precision, record.Recall, record.F1Score, record.Support)
}
