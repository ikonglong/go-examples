package optional_param

import "fmt"

type queryOptions struct {
	start string
	limit int
}

func find(table string, opts ...queryOptions) {
	fmt.Printf("opts: %v", opts)
}
