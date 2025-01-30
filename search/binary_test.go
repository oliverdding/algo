// Package search contains algorithms that related to searching.
package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	{
		position, shouldBePosition := BinarySearch([]int{1, 2, 3, 4, 5}, 1)
		assert.EqualValues(t, 0, position, "simple test")
		assert.EqualValues(t, -1, shouldBePosition, "simple test")
	}

	{
		position, shouldBePosition := BinarySearch([]int{1, 2, 4, 5}, 3)
		assert.EqualValues(t, -1, position, "simple test not found")
		assert.EqualValues(t, 2, shouldBePosition, "simple test not found")
	}

	{
		position, shouldBePosition := BinarySearch([]int{}, 1)
		assert.EqualValues(t, -1, position, "zero element test")
		assert.EqualValues(t, 0, shouldBePosition, "zero element test")
	}

	{
		position, shouldBePosition := BinarySearch([]int{5}, 5)
		assert.EqualValues(t, 0, position, "one element test (exist)")
		assert.EqualValues(t, -1, shouldBePosition, "one element test (exist)")
	}

	{
		position, shouldBePosition := BinarySearch([]int{5}, 0)
		assert.EqualValues(t, -1, position, "one element test (not exist, less)")
		assert.EqualValues(t, 0, shouldBePosition, "one element test (not exist, less)")
	}

	{
		position, shouldBePosition := BinarySearch([]int{5}, 6)
		assert.EqualValues(t, -1, position, "one element test (not exist, greater)")
		assert.EqualValues(t, 1, shouldBePosition, "one element test (not exist, greater)")
	}
}
