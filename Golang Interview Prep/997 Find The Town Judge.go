package l33tc0d3

// Solution comes from leetcode. I am commenting on it to make it easy to understand
func findJudge(N int, trust [][]int) int {

	// Base case. If no one trusts anyone and there's 1 person, he must be town judge.
	if len(trust) == 0 && N == 1 {
		return 1
	}

	// Make edges between the townspeople using the trust array.
	edges := make([][2]int, N+1)
	candidates := []int{}

	for _, t := range trust {
		edges[t[0]][0]++
		edges[t[1]][1]++

		if edges[t[1]][1] == N-1 { // if everyone in town trusts a person they may be the judge
			candidates = append(candidates, t[1])
		}
	}

	for _, c := range candidates {
		if edges[c][0] == 0 { // check if the candidate trusts no one
			return c
		}
	}

	return -1 // failed to find a town judge.
}

// more optimized solution - I did this problem as an intro to graphs so I'm not too concerned about this, but it's more elegant and simple
/*
func findJudge(N int, trust [][]int) int {
	data := make([]int, N+1)
	for _, t := range trust {
		data[t[0]], data[t[1]] = data[t[0]]-1, data[t[1]]+1
	}
	for i := 1; i <= N; i++ {
		if data[i] == N-1 {
			return i
		}
	}
	return -1
}
*/
