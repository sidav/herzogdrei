package game

import (
	"herzog/lib/random"
	"math"
)

var rnd random.PRNG

func SetPRNG(r random.PRNG) {
	rnd = r
}

func areFloatsRoughlyEqual(f, g float64) bool {
	return math.Abs(f-g) < 0.01
}

func abs(x int) int {
	// temp = value >> 31 // make a mask of the sign bit
	// value ^= temp      // toggle the bits if value is negative
	// value += temp & 1  // add one if value was negative
	if x < 0 {
		return -x
	}
	return x
}
