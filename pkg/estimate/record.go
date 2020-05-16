package estimate

import (
	"fmt"
)

func NewRecord(support int, precision, recall, f1Score float64) *Record {
	if support < 0 {
		panic(fmt.Sprintf("estimate: record got invalid support, %d", support))
	}
	if precision < 0 || precision > 1 {
		panic(fmt.Sprintf("estimate: record got invalid precision, %.3f", precision))
	}
	if recall < 0 || recall > 1 {
		panic(fmt.Sprintf("estimate: record got invalid recall, %.3f", recall))
	}
	if f1Score < 0 || f1Score > 1 {
		panic(fmt.Sprintf("estimate: record got invalid f1-score, %.3f", f1Score))
	}
	return &Record{support, precision, recall, f1Score}
}

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
