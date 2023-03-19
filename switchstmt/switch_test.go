package switchstmt

import (
	"fmt"
	"testing"
	"time"
)

// doc: https://yourbasic.org/golang/switch-statement/#basic-switch-with-default

// A switch statement runs the first case equal to the condition expression.
// The cases are evaluated from top to bottom, stopping when a case succeeds.
// If no case matches and there is a default case, its statements are executed.
func Test_BasicSwitchWithDefault(t *testing.T) {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Today is unknown")
	}
}

// First the switch expression is evaluated once.
// Then case expressions are evaluated left-to-right and top-to-bottom:
//     the first one that equals the switch expression triggers execution of the statements of the associated case,
// 	   the other cases are skipped.
func Test_ExecutionOrder(t *testing.T) {
	// Foo prints and returns n.
	foo := func(n int) int {
		fmt.Println(n)
		return n
	}

	switch foo(2) {
	// 先测试是否匹配 foo(1)，再测试是否匹配 foo(2)，然后 fallthrough 到下一个 case
	case foo(1), foo(2), foo(3):
		fmt.Println("First case")
		fallthrough
	case foo(4):
		fmt.Println("Second case")
	}
}
