// Package search contains algorithms that related to searching.
package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchLast(t *testing.T) {
	{
		position, shouldBePosition := BinarySearchLast([]int{1, 2, 3, 4, 5}, 1)
		assert.EqualValues(t, 0, position, "simple test")
		assert.EqualValues(t, -1, shouldBePosition, "simple test")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{1, 2, 4, 5}, 3)
		assert.EqualValues(t, -1, position, "simple test not found")
		assert.EqualValues(t, 2, shouldBePosition, "simple test not found")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{}, 1)
		assert.EqualValues(t, -1, position, "zero element test")
		assert.EqualValues(t, 0, shouldBePosition, "zero element test")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{5}, 5)
		assert.EqualValues(t, 0, position, "one element test (exist)")
		assert.EqualValues(t, -1, shouldBePosition, "one element test (exist)")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{5}, 0)
		assert.EqualValues(t, -1, position, "one element test (not exist, less)")
		assert.EqualValues(t, 0, shouldBePosition, "one element test (not exist, less)")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{5}, 6)
		assert.EqualValues(t, -1, position, "one element test (not exist, greater)")
		assert.EqualValues(t, 1, shouldBePosition, "one element test (not exist, greater)")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{1, 1, 1}, 1)
		assert.EqualValues(t, 2, position, "duplicate elements test")
		assert.EqualValues(t, -1, shouldBePosition, "duplicate elements test")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{1, 1, 1}, 0)
		assert.EqualValues(t, -1, position, "duplicate elements test (not exist, less)")
		assert.EqualValues(t, 0, shouldBePosition, "duplicate elements test (not exist, less)")
	}

	{
		position, shouldBePosition := BinarySearchLast([]int{1, 1, 1}, 2)
		assert.EqualValues(t, -1, position, "duplicate elements test (not exist, greater)")
		assert.EqualValues(t, 3, shouldBePosition, "duplicate elements test (not exist, greater)")
	}
}
