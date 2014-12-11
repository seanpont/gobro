package strarr

import (
	_ "fmt"
	"github.com/seanpont/assert"
	"testing"
)

func TestTrimAll(t *testing.T) {
	assert := assert.Assert(t)
	items := []string{"  one ", "\ttwo ", "three\n\n"}
	TrimAll(items)
	assert.Equal(items[0], "one")
	assert.Equal(items[1], "two")
	assert.Equal(items[2], "three")
}

func TestIndexOf(t *testing.T) {
	assert := assert.Assert(t)
	// encryption works
	items := []string{"one", "two", "three"}
	assert.Equal(IndexOf(items, "one"), 0)
	assert.Equal(IndexOf(items, "four"), -1)
}

func TestContains(t *testing.T) {
	assert := assert.Assert(t)
	items := []string{"one", "two", "three"}
	assert.True(Contains(items, "one"), "one")
	assert.False(Contains(items, "four"), "four")
}

func TestFindMatchWithRegex(t *testing.T) {
	assert := assert.Assert(t)
	items := []string{"one", "two.2", "thr33"}
	one, err := FindMatchWithRegex(items, "on\\w")
	assert.Nil(err)
	assert.Equal(one, "one")
	three, err := FindMatchWithRegex(items, ".*\\d\\d")
	assert.Nil(err)
	assert.Equal(three, "thr33")
	what, err := FindMatchWithRegex(items, "f.*")
	assert.Nil(err)
	assert.Equal(what, "")
	bad, err := FindMatchWithRegex(items, "\\q")
	assert.NotNil(err)
	assert.Equal(bad, "")
}
