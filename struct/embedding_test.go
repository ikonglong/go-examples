package struct_ex

import (
	"fmt"
	"github.com/ikonglong/go-examples/struct/embedding"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

type Animal struct {
	name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating\n", a.name)
}

type Dog struct {
	Animal
}

func TestCreateSubType(t *testing.T) {
	dog := Dog{
		Animal{
			name: "Snoopy",
		},
	}
	dog.Eat()
	fmt.Printf("dog name: %v\n", dog.name)
}

func TestEmbedding(t *testing.T) {
	convey.Convey("假设类型 A 嵌入了类型 B，A、B 在同一个包，且 B 未导出", t, func() {
		convey.Convey("当尝试访问 A 对象中内嵌的 B 的导出成员时", func() {
			u := embedding.User{Name: "bob"}
			u.City = "shanghai"
			convey.Convey("那么访问成功", func() {
				convey.So(u.City, convey.ShouldEqual, "shanghai")
			})
		})
	})
}
