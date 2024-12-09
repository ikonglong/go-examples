package mapusage

import (
	"testing"
)

func TestStringKey(t *testing.T) {
	aMap := make(map[string]int)
	keyA1 := "a"
	keyA2 := "a"
	t.Logf("&keyA1: %p, &keyA2: %p", &keyA1, &keyA2)
	aMap[keyA1] = 1
	t.Logf("%v: %v", keyA2, aMap[keyA2])
}

func TestStringPtrKey(t *testing.T) {
	aMap := make(map[*string]int)
	key1 := "a"
	key2 := "a"
	t.Logf("&key1: %p, &key2: %p", &key1, &key2)

	keyA1Ptr := &key1
	aMap[keyA1Ptr] = 1

	v, present := aMap[&key2]
	if present {
		t.Logf("aMap[&key2]: %v", v)
	} else {
		t.Logf("aMap[&key2] not present")
	}

	key3 := &key1
	t.Logf("key3 point to key1, aMap[key3]: %v", aMap[key3])

	// Error: Cannot use '*key3' (type string) as the type *string
	//t.Logf("*key3[=%v]: %v", *key3, aMap[*key3])
	// Error: Cannot use '&key3' (type **string) as the type *string
	//t.Logf("&key3[=%v]: %v", *key3, aMap[&key3])
}

type BookTitle struct {
	title string
}

type BookTitlePtr struct {
	titlePtr *string
}

func TestBookTitleKey(t *testing.T) {
	aMap := make(map[BookTitle]string)
	title1 := BookTitle{
		title: "Go",
	}
	aMap[title1] = "content"
	title2 := BookTitle{
		title: "Go",
	}
	t.Logf("&title1: %p, &title2: %p", &title1, &title2)
	t.Logf("aMap[title2]: %v\n", aMap[title2])
}

func TestBookTitlePtrKey(t *testing.T) {
	titleVal := "Go"
	titleValPtr1 := &titleVal
	titleValPtr2 := &titleVal

	aMap := make(map[BookTitlePtr]string)
	title1 := BookTitlePtr{
		titlePtr: titleValPtr1,
	}
	aMap[title1] = "content"

	title2 := BookTitlePtr{
		titlePtr: titleValPtr2,
	}

	// title1.titlePtr 和 title2.titlePtr 这两个指针变量的值相同，所以以 title2 为键，也能取到值
	t.Logf("title2: %v\n", aMap[title2])
}
