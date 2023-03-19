package error

import (
	"fmt"
	"testing"
)

func TestIllegalArg(t *testing.T) {
	err := doSth2()
	if err == nil {
		fmt.Println("err == nil")
	}
	fmt.Println("---------------")
	//handleError(err)
}
