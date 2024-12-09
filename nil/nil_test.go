package nil

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroMapIsNil(t *testing.T) {
	var m map[string]interface{}
	assert.True(t, m == nil)
}

type domainError struct{}

func (e domainError) Error() string {
	return "some domain error"
}

func TestCheckIsNilAfterTypeAssertion(t *testing.T) {
	convey.Convey("Given func f returns nil of Type ", t, func() {
		f := func() error {
			var e *domainError = nil
			return e
		}
		convey.Convey("When f is called, r received the return val,"+
			" get the interface value's underlying concrete value with type assertion", func() {
			r := f()
			v := r.(*domainError)
			convey.Convey("Then r is not nil, v is nil", func() {
				convey.So(r != nil, convey.ShouldEqual, true)
				convey.So(v == nil, convey.ShouldEqual, true)
			})
		})
	})
}
