package convey

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestXxx(t *testing.T) {
	// the setup logic which should be run once before all the BDD-style tests here run starts

	// setup ...

	// the setup logic which should be run once before all the BDD-style tests here run ends

	Convey("Given ...", t, func() {
		// ...
		Convey("When ...", func() {
			// ...
			Convey("Then ...", func() {
				// ...
				// So(err, ShouldBeNil)
				// So(c, ShouldBeNil)
			})
		})

		Convey("When ...", func() {
			// ...
			Convey("Then ...", func() {
				// ...
				// So(err, ShouldBeNil)
				// So(c, ShouldBeNil)
			})
		})

		Convey("When ...", func() {
			// ...
			Convey("Then ...", func() {
				// ...
				// So(err, ShouldBeNil)
				// So(c, ShouldBeNil)
			})
		})
	})
}
