package costs

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewMSECostFunction() *MSECostFunction {
	return &MSECostFunction{}
}

type MSECostFunction struct{}

func (costFunction *MSECostFunction) Evaluate(actual mat.Vector) int {
	if actual == nil {
		panic("costs: mse cost function got nil actual")
	}
	index, max := 0, actual.AtVec(0)
	for i := 1; i < actual.Len(); i++ {
		if value := actual.AtVec(i); value > max {
			index = i
			max = value
		}
	}
	return index
}

func (costFunction *MSECostFunction) Differentiate(actual mat.Vector, label int) mat.Vector {
	if actual == nil {
		panic("costs: mse cost function got nil actual")
	}
	length := actual.Len()
	if label < 0 || label >= length {
		panic(fmt.Sprintf("costs: mse cost function got invalid label, %d", label))
	}
	costs := mat.NewVecDense(length, nil)
	costs.SetVec(label, 1)
	costs.SubVec(actual, costs)
	return costs
}
