package func_value

import "sort"

type Operator func(x float64) float64

// https://yourbasic.org/golang/function-pointer-type-declaration/
// Map applies op to each element of a.
func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}

// https://yourbasic.org/golang/anonymous-function-literal-lambda-closure/
func SortSlice(slice interface{}, less func(i, j int) bool) {
	sort.Slice(slice, less)
}
