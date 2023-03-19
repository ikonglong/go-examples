package optional_param

import "testing"

func TestFind(t *testing.T) {
	x := []queryOptions{
		{
			start: "start_id",
			limit: 50,
		},
	}
	find("books", x...)
}
