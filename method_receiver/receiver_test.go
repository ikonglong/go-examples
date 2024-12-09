package method_receiver

import (
	"fmt"
	"testing"
)

type opError struct {
	id      string
	details map[string]any
}

func (e opError) WithDetail(key string, value any) {
	e.details[key] = value
	fmt.Printf("copy of this opError: %v\n", e)
}

func (e opError) WithDetail2(key string, value any) *opError {
	copy := &opError{
		id:      e.id,
		details: e.details,
	}
	copy.details[key] = value
	fmt.Printf("copy of this opError: %v\n", copy)
	return copy
}

var internalErrPrototype = opError{
	id:      "internal error",
	details: map[string]any{},
}

func TestValueTypeReceiver(t *testing.T) {
	fmt.Printf("original internalErr: %v\n", internalErrPrototype)
	internalErrPrototype.WithDetail("x1", "x1_v")
	fmt.Printf("original internalErr: %v\n", internalErrPrototype)

	fmt.Printf("==========================\n")

	fmt.Printf("original internalErr: %v\n", internalErrPrototype)
	internalErrPrototype.WithDetail("x2", "x2_v")
	fmt.Printf("original internalErr: %v\n", internalErrPrototype)
}
