package switchstmt

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

// doc: https://yourbasic.org/golang/switch-statement/

func TestSwitch1(t *testing.T) {
	// the setup logic which should be run once before all the BDD-style tests here run starts

	// ...

	// the setup logic which should be run once before all the BDD-style tests here run ends

	convey.Convey("假设有一 switch 块，参数跟最后一个 case condition 匹配", t, func() {
		arg := time.Saturday
		sayWeekday := func(theDay time.Weekday) {
			switch theDay {
			case time.Sunday:
				fmt.Println("Today is Sunday")
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
			default:
				fmt.Println("Today is unknown")
			}
		}
		convey.Convey("当将此参数传递给给 switch 块", func() {
			sayWeekday(arg)
			convey.Convey("那么每个 case 的条件匹配都会执行，且只会执行最后一个 case 中的语句", func() {
				// do nothing
				//convey.So(true, convey.ShouldBeTrue)
			})
		})
	})

}

func TestSwitch2(t *testing.T) {
	convey.Convey("假设有一 switch 块，参数跟中间某个 case 匹配，"+
		"且此 case 块内部没有 break/return，且跟后续的每个 case 都匹配（都没有 break/return）", t, func() {
		arg := time.Wednesday
		toInt := func(wd time.Weekday) int {
			fmt.Printf("wd to int: %v\n", wd)
			return int(wd)
		}
		sayWeekday := func(theDay time.Weekday) {
			switch v := int(theDay); {
			case v <= toInt(time.Sunday):
				fmt.Println("Until Sunday")
			case v <= toInt(time.Monday):
				fmt.Println("Until Monday")
			case v <= toInt(time.Tuesday):
				fmt.Println("Until Tuesday")
			case v <= toInt(time.Wednesday):
				fmt.Println("Until Wednesday")
			case v <= toInt(time.Thursday):
				fmt.Println("Until Thursday")
			case v <= toInt(time.Friday):
				fmt.Println("Until Friday")
			case v <= toInt(time.Saturday):
				fmt.Println("Until Saturday")
			default:
				fmt.Println("Today is unknown")
			}
		}
		convey.Convey("当将此参数传递给此 switch 块", func() {
			sayWeekday(arg)
			convey.Convey("那么只会执行第一个匹配的 case，后续能够匹配的 case 都不会执行", func() {
				// A switch statement runs the first case equal to the condition expression.
				// The cases are evaluated from top to bottom, stopping when a case succeeds.
				// If no case matches and there is a default case, its statements are executed.
			})
		})
	})
}

// First the switch expression is evaluated once.
// Then case expressions are evaluated left-to-right and top-to-bottom:
//
//	    the first one that equals the switch expression triggers execution of the statements of the associated case,
//		   the other cases are skipped.
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
