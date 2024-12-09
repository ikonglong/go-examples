package mapusage

import (
	"fmt"
	"testing"
)

func TestInterfaceValuePointToInt(t *testing.T) {
	aMap := make(map[string]interface{})
	key := "a"
	originalVal := 10
	// 因为 int 是值类型，所以这里赋值操作的右值是 originalVal 的副本
	aMap[key] = originalVal
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// 因为 map 中的 value 指针指向的是 originalVal 的副本，所以修改 originalVal，map 中的值并不会改变
	originalVal = 20
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// map 中的 value 指针现在指向了值为 30 的新变量
	aMap[key] = 30
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])
}

func TestInterfaceValuePointToString(t *testing.T) {
	aMap := make(map[string]interface{})
	key := "a"
	originalVal := "va"
	// 因为 string 是值类型，所以这里赋值操作的右值是 originalVal 的副本
	aMap[key] = originalVal
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// 因为 map 中的 value 指针指向的是 originalVal 的副本，所以修改 originalVal，map 中的值并不会改变
	originalVal = "vb"
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// map 中的 value 指针现在指向了值为 'vc' 的新变量
	aMap[key] = "vc"
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])
}

func TestInterfaceValuePointToStringPtr(t *testing.T) {
	aMap := make(map[string]interface{})
	key := "a"
	originalVal := "va"
	originalValPtr := &originalVal
	aMap[key] = originalValPtr
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, *(aMap[key].(*string)))

	originalVal = "vb"
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, *(aMap[key].(*string)))

	// map 中的 value 指针现在指向了值为 'vc' 的新变量
	newVal := "vc"
	aMap[key] = &newVal
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, *(aMap[key].(*string)))
}

func TestInterfaceValuePointToStruct(t *testing.T) {
	type book struct {
		title string
	}

	aMap := make(map[string]interface{})
	key := "a"
	originalVal := book{
		title: "title_a",
	}
	// 因为 struct 是值类型，所以这里赋值操作的右值是 originalVal 的副本
	aMap[key] = originalVal
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	originalVal.title = "title_b"
	// 因为 map 中的 value 指针指向的是 originalVal 的副本，所以修改 originalVal，map 中的值并不会改变
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	originalVal = book{
		title: "title_c",
	}
	// 因为 map 中的 value 指针指向的是 originalVal 的副本，所以修改 originalVal，map 中的值并不会改变
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// map 中的 value 指针现在指向了一个新的结构体变量
	aMap[key] = book{
		title: "title_d",
	}
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])
}

func TestInterfaceValPointToSlice(t *testing.T) {
	originalVal := make([]string, 0, 5)
	// 如果使用 _ 接收 append 的返回值，originalVal 仍为空列表。
	// 猜测是因为 append 返回了一个新的标头结构体实例来表示更新后的切片
	originalVal = append(originalVal, "a")
	originalVal = append(originalVal, "b")
	originalVal = append(originalVal, "c")

	aMap := make(map[string][]string)
	key := "a"
	// 因为 slice 是值类型，所以这里赋值操作的右值是 originalVal 的副本
	aMap[key] = originalVal
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// append 应该是返回了一个新的标头实例来表示追加后的切片。但 map 中的 value 指针
	// 指向的是追加前标头实例的副本（即只包含 a,b,c），因此 map 中的值不会改变
	originalVal = append(originalVal, "d")
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// 这一步追加前，map 中的 value 指针指向的是前一步追加前标头实例的副本（即只包含 a,b,c）。
	// 将 map 中的 value 指针指向了追加后返回的新标头值后，map 中的 slice 包含 a,b,c,e。
	// 在执行该步之前，originalVal 包含 a,b,c,d。执行该步之后，由于共享底层数组，元素 d 被
	// e 覆盖，因此 originalVal 也包含 a,b,c,e。
	aMap[key] = append(aMap[key], "e")
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])
}

func TestInterfaceValPointToMap(t *testing.T) {
	originalVal := map[string]string{
		"name": "name_a",
	}

	aMap := make(map[string]interface{})
	key := "a"
	aMap[key] = originalVal
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	originalVal["name"] = "name_b"
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// Cannot convert an expression of the type 'interface{}' to the type 'map[string]string'
	//valMap := map[string]string(aMap[key])
	valMap := aMap[key].(map[string]string)
	valMap["name"] = "name_c"
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])

	// map 不是基础类型，因此这里其实是将 originalVal 指向了一个新的 map 对象
	originalVal = map[string]string{
		"name": "name_d",
	}
	fmt.Printf("originalVal: %v, aMap[key]: %v\n", originalVal, aMap[key])
}

func TestXxx(t *testing.T) {
	type key struct {
		name string
	}
	aMap := make(map[key][]int, 8)
	k1 := key{
		name: "k1",
	}
	aMap[k1] = append(aMap[k1], 0)
	fmt.Printf("aMap: %v\n", aMap)

	v2 := aMap[key{name: "k2"}]
	fmt.Printf("value for 'k2' is nil: %v\n", v2 == nil)
	fmt.Printf("value for 'k2': %v\n", v2)
}
