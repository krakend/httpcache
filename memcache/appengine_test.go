//go:build appengine
// +build appengine

package memcache

import (
	"appengine/aetest"
	"testing"

	"github.com/gregjones/httpcache/test"
)

func TestAppEngine(t *testing.T) {
	ctx, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	test.Cache(t, New(ctx))
}
