package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_Add(t *testing.T) {
	is := assert.New(t)

	set := New[int](3)

	set.Add(1)

	_, ok := set.m[1]
	is.True(ok)
}

func TestSet_Remove(t *testing.T) {
	is := assert.New(t)

	set := New[int](3)

	set.Add(1)
	set.Remove(1)

	_, ok := set.m[1]
	is.False(ok)
}

func TestSet_Contains(t *testing.T) {
	is := assert.New(t)

	set := New[int](3)

	set.Add(1)

	actual := set.Contains(1)
	is.True(actual)
}

func TestSet_Size(t *testing.T) {
	is := assert.New(t)

	set := New[int](5)

	set.Add(1)
	set.Add(2)
	set.Add(3)

	expected := 3
	actual := set.Size()
	is.EqualValues(expected, actual)
}
