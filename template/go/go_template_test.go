package gotempl

import (
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"html/template"
	"strings"
	"testing"
)

func Create(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}

func TestBasicUsage(t *testing.T) {
	tpl := template.New("t")
	// tpl, err := tpl.Parse("value is {{.}}")
	// if err != nil {
	// 	panic(err)
	// }
	tpl = template.Must(tpl.Parse("value is {{.}}"))

	var sb = &strings.Builder{}
	tpl.Execute(sb, "hello")
	assert.Equal(t, "value is hello", sb.String())

	sb = &strings.Builder{}
	tpl.Execute(sb, 1)
	assert.Equal(t, "value is 1", sb.String())

	sb = &strings.Builder{}
	tpl.Execute(sb, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
	assert.Equal(t, "value is [Go Rust C&#43;&#43; C#]", sb.String())

	tpl = Create("t", "value is {{.Name}}")
	sb.Reset()
	tpl.Execute(sb, struct {
		Name string
	}{"Jane"})
	assert.Equal(t, "value is Jane", sb.String())
	sb.Reset()
	tpl.Execute(sb, map[string]string{
		"Name": "Bob",
	})
	assert.Equal(t, "value is Bob", sb.String())

	tpl = Create("t", "value is {{.name}}")
	sb.Reset()
	tpl.Execute(sb, struct {
		name string
	}{"Jane"})
	assert.Equal(t, "value is ", sb.String())
	sb.Reset()
	tpl.Execute(sb, map[string]string{
		"name": "Bob",
	})
	assert.Equal(t, "value is Bob", sb.String())

	tpl = Create("tpl",
		// - 将邻接的空白字符去掉
		"{{if . -}} yes {{ else -}} no {{end}}")
	sb.Reset()
	tpl.Execute(sb, "not empty")
	assert.Equal(t, "yes ", sb.String())
	sb.Reset()
	tpl.Execute(sb, "")
	assert.Equal(t, "no ", sb.String())

	tpl = Create("t",
		"list: {{range .}}{{.}} {{end}}")
	sb.Reset()
	tpl.Execute(sb, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
	assert.Equal(t, "list: Go Rust C&#43;&#43; C# ", sb.String())

	tpl = Create("t",
		"map: {{range $key, $value := .}}{{$key}}:{{$value}},{{end}}")
	sb.Reset()
	tpl.Execute(sb, map[string]interface{}{
		"a": 1,
		"b": 2,
	})
	assert.Equal(t, "map: a:1,b:2,", sb.String())

	// template text: "map: {{range $key, $value := users}}{{$key}}:{{$value}},{{end}}", panic
	tpl = Create("t",
		"map: {{range $key, $value := .users}}{{$key}}:{{$value}},{{end}}")
	sb.Reset()
	tpl.Execute(sb, map[string]interface{}{
		"users": map[string]interface{}{
			"alice": 20,
			"bob":   25,
		},
	})
	assert.Equal(t, "map: alice:20,bob:25,", sb.String())
}

func TestIsNull(t *testing.T) {
	// test the root data object is not null
	tpl := Create("tpl",
		// - 将邻接的空白字符去掉
		"{{if . -}} not null {{- else -}} null {{- end}}")
	// `...{{-end}}` 报告错误：panic: template: tpl:1: bad number syntax: "-en"

	var sb = &strings.Builder{}
	tpl.Execute(sb, nil)
	assert.Equal(t, "null", sb.String())

	var dataMap map[string]interface{} = nil
	assert.True(t, dataMap == nil)
	var data any = dataMap
	assert.False(t, data == nil)

	sb.Reset()
	tpl.Execute(sb, data)
	assert.Equal(t, "null", sb.String())

	convey.Convey("假设模版使用 `if ne . nil` 测试数据对象是否不为 nil，且数据对象为 nil", t, func() {
		tplStr := "{{if ne . nil}} not null {{- else -}} null {{- end}}"
		convey.Convey("当将此数据对象应用于模版时", func() {
			tpl := Create("tpl", tplStr)
			sb.Reset()
			tpl.Execute(sb, nil)

			convey.Convey("那么应该应用成功，且结果为 null", func() {
				convey.So(sb.String(), should.Equal, "null")
			})
		})
	})

	convey.Convey("假设模版使用 `if ne . null` 测试数据对象是否为 nil，且数据对象为 nil", t, func() {
		tplStr := "{{if ne . null}} not null {{- else -}} null {{- end}}"
		convey.Convey("当将此数据对象应用于模版时", func() {
			defer func() {
				r := recover()
				convey.Convey("那么应该报告错误", func() {
					convey.So(r, convey.ShouldNotBeNil)
					convey.So(r, convey.ShouldBeError)
					err := r.(error)
					convey.So(err.Error(), convey.ShouldContainSubstring, "function \"null\" not defined")
				})
			}()

			tpl := Create("tpl", tplStr)
			sb.Reset()
			tpl.Execute(sb, nil)
		})
	})

	convey.Convey("假设模版使用看起来非法的 `if ne .` 测试数据对象是否为 nil，且数据对象为 nil", t, func() {
		tplStr := "{{if ne .}} not null {{- else -}} null {{- end}}"
		convey.Convey("当将此数据对象应用于模版时", func() {
			tpl := Create("tpl", tplStr)
			sb.Reset()
			tpl.Execute(sb, nil)
			convey.Convey("那么应该应用成功，且结果为 \"\"，虽然结果出人意料", func() {
				convey.So(sb.String(), should.Equal, "")
			})
		})
	})

	convey.Convey("假设模版使用看起来非法的 `if ne nil` 测试数据对象是否为 nil，且数据对象为 nil", t, func() {
		tplStr := "{{if ne nil}} not null {{- else -}} null {{- end}}"
		convey.Convey("当将此数据对象应用于模版时", func() {
			tpl := Create("tpl", tplStr)
			sb.Reset()
			tpl.Execute(sb, nil)
			convey.Convey("那么应该应用成功，且结果为 \"\"，虽然结果出人意料", func() {
				convey.So(sb.String(), should.Equal, "")
			})
		})
	})
}
