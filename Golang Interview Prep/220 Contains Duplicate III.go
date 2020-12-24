package l33tc0d3

import (
	"fmt"
	"math"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// best case :)
	if len(nums) < 2 || k <= 0 {
		return false
	}

	// otherwise have to iterate
	width := t + 1
	buckets := make(map[int]int)
	buckIdx := 0
	for i, n := range nums {
		// bucket placement edge case
		if n < 0 {
			buckIdx = int(math.Floor(float64(n) / float64(width)))
		} else {
			buckIdx = n / width
		}
		if _, ok := buckets[buckIdx]; ok { // current index has a value already, aka within range
			return true
		}

		buckets[buckIdx] = n
		// out of bucket range check
		if _, ok := buckets[buckIdx-1]; ok {
			if math.Abs(float64(buckets[buckIdx-1]-n)) <= math.Abs(float64(t)) {
				return true
			}
		}
		if _, ok := buckets[buckIdx+1]; ok {
			if math.Abs(float64(buckets[buckIdx+1]-n)) <= math.Abs(float64(t)) {
				return true
			}
		}

		// slide window update check
		fmt.Println("i: ", i)
		if i >= k {
			fmt.Println("deleting index ", nums[i-k]/width)
			delete(buckets, nums[i-k]/width)
		}

	}

	// condition not met, return false
	return false
}
