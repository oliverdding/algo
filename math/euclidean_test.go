// Package math contains algorithms that related to math.
package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuclideanGCD(t *testing.T) {
	assert.EqualValues(t, 5, EuclideanGCD(0, 5), "the first parameter is 0")
	assert.EqualValues(t, 105, EuclideanGCD(105, 0), "the second parameter is 0")
	assert.EqualValues(t, 0, EuclideanGCD(0, 0), "all parameters are 0")
	assert.EqualValues(t, 6, EuclideanGCD(24, 18), "gcd(24, 18) should be 6")
	assert.EqualValues(t, 5, EuclideanGCD(20, 15), "gcd(20, 15) should be 5")
}
