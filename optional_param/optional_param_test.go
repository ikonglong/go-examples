package optional_param

import (
	"fmt"
	"testing"
)

type queryOptions struct {
	start string
	limit int
}

func find(table string, opts ...queryOptions) {
	fmt.Printf("opts: %v", opts)
}

func TestFind(t *testing.T) {
	x := []queryOptions{
		{
			start: "start_id",
			limit: 50,
		},
	}
	find("books", x...)
}

func TestXxx(t *testing.T) {
	var print = func(values ...string) {
		for _, v := range values {
			fmt.Printf("%s, ", v)
		}
	}
	var delegatePrint = func(values ...string) {
		print(values...)
	}
	delegatePrint("a", "b", "c")
	delegatePrint([]string{"a", "b", "c"}...)
}
