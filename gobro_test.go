package gobro

import (
	_ "fmt"
	"github.com/seanpont/assert"
	"testing"
)

func TestIndexOf(t *testing.T) {
	assert := assert.Assert(t)
	// encryption works
	items := []string{"one", "two", "three"}
	assert.Equal(IndexOf(items, "one"), 0)
	assert.Equal(IndexOf(items, "four"), -1)
}
