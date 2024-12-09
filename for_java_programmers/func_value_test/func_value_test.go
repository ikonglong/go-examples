package func_value_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
	"github.com/ikonglong/go-examples/for_java_programmers/func_value"
)

func TestFuncValue(t *testing.T) {
	op := math.Abs
	a := []float64{1, -2}
	b := func_value.Map(op, a)
	fmt.Println(b)

	c := func_value.Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c)
}

func TestSortSlice(t *testing.T) {
	people := []string{"Alice", "Bob", "Dave"}
	// people := [...]string{"Alice", "Bob", "Dave"} // runtime error
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})

	// You can also use an intermediate variable.
	less := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}
	sort.Slice(people, less)
	// Note that the less function is a closure: it references the people variable,
	// which is declared outside the function.

	fmt.Println(people)
	// Output: [Bob Dave Alice]
}
