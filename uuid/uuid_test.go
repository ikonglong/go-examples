package uuid

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 3; i++ {
		uuid := uuid2.New()
		fmt.Printf("%v\n", uuid.String())
	}
}
