package l33tc0d3

import "math"

func rangeSum(nums []int, n int, left int, right int) int {
	// method: counting - more efficient than sorting the numbers. not the intuitive approach
	// but easy to understand and more performant
	// full credit to magicwing on leetcode for code

	// first get the total sum of all numbers

	// then loop and count all the partial sums from index 0 to to n-1, making a slice counting each index
	// loop and count all partial sums again from index 1 to n-1, then 2 to n-1, etc.

	// then loop over all numbers from 1 to sum to find if anything within range is a partial sum.
	// if it is, and it is within range, add to return value

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	arr := make([]int, sum+1)
	for i := 0; i < n; i++ {
		amount := 0
		for j := i; j < n; j++ {
			amount += nums[j]
			arr[amount]++
		}
	}
	ret, idx := 0, 0
	for i := 1; i < sum+1; {
		if arr[i] > 0 {
			arr[i]--
			idx++
			if idx >= left && idx <= right {
				ret += i
			} else if idx > right {
				break
			}
		} else {
			i++
		}
	}
	return ret % int(math.Pow10(9)+7)
}
