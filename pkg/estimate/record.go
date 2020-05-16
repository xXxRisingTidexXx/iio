package estimate

import (
	"fmt"
)

type Record struct {
	Support   int
	Precision float64
	Recall    float64
	F1Score   float64
}

func (record *Record) Equal(other *Record) bool {
	return other != nil &&
		record.Support == other.Support &&
		record.Precision == other.Precision &&
		record.Recall == other.Recall &&
		record.F1Score == other.F1Score
}

func (record *Record) String() string {
	return fmt.Sprintf("{%d %f %f %f}", record.Support, record.Precision, record.Recall, record.F1Score)
}
