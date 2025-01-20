package test_test

import (
	"testing"

	"github.com/krakendio/httpcache"
	"github.com/krakendio/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}

func TestLruCache(t *testing.T) {
	test.Cache(t, httpcache.NewLruCache(10))
}
