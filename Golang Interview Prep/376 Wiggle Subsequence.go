package l33tc0d3

// So apparently the brute force method i did sucks and just evaluates the same as the greedy approach
// ...so ummm just do greedy algorithm for best result here.

// Attempt 1 - brute force
/*
func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	longestSS := 1  // longest subsequence initial value
	ssBuilder := 1  // temporary value that will update longest_ss
	prev := nums[0] // initial value
	diff := 0       // expected diff may be either negative or positive. 0 to start with before decision is made
	for i := range nums {
		ssBuilder = 1

		for j := i; j < len(nums); j++ {
			n := nums[j]

			if n > prev {
				if diff == 0 || diff == -1 {
					ssBuilder++
					diff = 1
				}
			} else if n < prev {
				if diff == 0 || diff == 1 {
					ssBuilder++
					diff = -1
				}
			}
			prev = n
		}
		if ssBuilder > longestSS {
			longestSS = ssBuilder
		}
	}
	return longestSS
}
*/

// Attempt 2 - greedy
func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	longestSS := 1 // longest subsequence initial value
	diff := 0      // expected diff may be either negative or positive. 0 to start with before decision is made
	for i, n := range nums {
		if i != 0 {
			if n < nums[i-1] {
				if diff == 0 || diff == 1 {
					longestSS++
					diff = -1
				}
			} else if n > nums[i-1] {
				if diff == 0 || diff == -1 {
					longestSS++
					diff = 1
				}
			} else { // n == nums[i-1]
				// DO NOTHING!
			}
		}
	}
	return longestSS
}
