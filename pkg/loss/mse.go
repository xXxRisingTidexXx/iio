package loss

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type MSECostFunction struct{}

func (costFunction *MSECostFunction) Evaluate(actual mat.Vector) int {
	if actual == nil {
		panic("loss: mse cost function got nil vector")
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
		panic("loss: mse cost function got nil vector")
	}
	length := actual.Len()
	if label < 0 || label >= length {
		panic(fmt.Sprintf("loss: mse cost function got invalid label, %d", label))
	}
	costs := mat.NewVecDense(length, nil)
	costs.SetVec(label, 1)
	costs.SubVec(actual, costs)
	return costs
}
