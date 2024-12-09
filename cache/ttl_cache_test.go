package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/stretchr/testify/assert"
)

// https://github.com/jellydator/ttlcache
func TestTTLCache(t *testing.T) {
	type Book struct {
		title string
	}

	cache := ttlcache.New[string, Book](
		ttlcache.WithTTL[string, Book](30 * time.Minute),
	)

	bookA := Book{
		title: "book a",
	}
	cache.Set("book_a", bookA, 10*time.Second)
	v := cache.Get("book_a").Value()
	fmt.Printf("origin value: %p, cached value: %p", &bookA, &v)
	assert.Fail(t, "")
}
