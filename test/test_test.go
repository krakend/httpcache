package test_test

import (
	"testing"

	"github.com/krakend/httpcache"
	"github.com/krakend/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
