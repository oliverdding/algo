// Package math contains algorithms that related to math.
package math

// EuclideanGCD return the greatest common divisor of two parameters using euclidean algorithm.
func EuclideanGCD(a, b uint) uint {
	if b == 0 {
		return a
	}

	return EuclideanGCD(b, a%b)
}
