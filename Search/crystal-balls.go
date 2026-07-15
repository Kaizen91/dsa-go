package search

import (
	"math"
)

// algorithmn to solve this is to jump in sqrt(n) increments
// then once found jump back to the previous solution and 
// linearly walk from the last point searching for the first
// true

func FindBreak(arr []bool) int {
	l := len(arr)
	if l == 0 {
		return -1
	}
	jump := int(math.Floor(math.Sqrt(float64(len(arr)))))
	prev := 0
	i := 0
	for i < l && !arr[i] {
		prev = i
		i += jump
	}

	//make sure we don't over jump
	if i > l {
		i = l - 1
	}
	
	for j := prev; j <= i; j++ {
		if arr[j] {
			return j
		}
	}
	return -1 
}
