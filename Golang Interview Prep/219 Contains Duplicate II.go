package l33tc0d3

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, n := range nums {
		if _, ok := m[n]; ok && math.Abs(float64(m[n]-i)) <= float64(k) {
			return true
		}
		m[n] = i
	}
	return false
}
