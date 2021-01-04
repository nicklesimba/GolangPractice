package l33tc0d3

// naive approach, apparently. this took 228ms, and used 22mb of memory.
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	// go through list of edges
	// map TO value to being found or not, ex: (0,1) -> to[1] = true (otherwise maintain default false)

	to := make([]bool, n)
	result := []int{}
	for _, n := range edges {
		to[n[1]] = true
	}

	// go through all n vertices
	// if from[i] > 0 and to[i] == 0, add i to the result
	for i, v := range to {
		if v != true {
			result = append(result, i)
		}
	}

	// return result
	return result
}
