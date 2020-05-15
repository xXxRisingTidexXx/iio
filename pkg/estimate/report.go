package estimate

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
)

type Report struct {
	Classes  []*Record
	MacroAvg *Record
	Accuracy float64
}

func (report *Report) Equal(other *Report) bool {
	return other != nil &&
		cmp.Equal(report.Classes, other.Classes) &&
		cmp.Equal(report.MacroAvg, other.MacroAvg) &&
		report.Accuracy == other.Accuracy
}

func (report *Report) String() string {
	builder := strings.Builder{}
	builder.WriteString("           | precision | recall | f1-score | support \n")
	builder.WriteString("-----------+-----------+--------+----------+---------\n")
	for i, class := range report.Classes {
		builder.WriteString(
			fmt.Sprintf(
				" class %3d | %9.3f | %6.3f | %8.3f | %7d \n",
				i,
				class.Precision,
				class.Recall,
				class.F1Score,
				class.Support,
			),
		)
	}
	builder.WriteString(
		fmt.Sprintf(
			" macro avg | %9.3f | %6.3f | %8.3f | %7d \n",
			report.MacroAvg.Precision,
			report.MacroAvg.Recall,
			report.MacroAvg.F1Score,
			report.MacroAvg.Support,
		),
	)
	builder.WriteString(
		fmt.Sprintf("  accuracy |           |        | %8.3f |         \n", report.Accuracy),
	)
	builder.WriteString("-----------+-----------+--------+----------+---------\n")
	return builder.String()
}
