package l33tc0d3

func allPathsSourceTarget(graph [][]int) [][]int {
	// Base case #1. If graph is empty return empty list
	if len(graph) == 0 || len(graph) == 1 {
		return [][]int{}
	}

	// Start from node 0, check its list and traverse in order.
	// If it reaches node n-1, add it to list and go to previous node,
	// check all its previous options. once all options are exhausted to previous node, etc.
	// Once node 0 has all its options exhausted, return resulting set.
	source, dest := 0, len(graph)-1
	paths := [][]int{}
	path := []int{}

	// messing with anonymous functions
	var dfs func(a int)

	dfs = func(vert int) {
		path = append(path, vert)

		if vert == dest { // destination node found!
			// we use a temp to not "contaminate" slice.
			// https://leetcode.com/problems/all-paths-from-source-to-target/discuss/887894/Go-Backtracking-beats-100-runtime
			// I'm pretty confused on why the temp is needed still.
			temp := make([]int, len(path))
			copy(temp, path)
			paths = append(paths, temp)
		} else {
			for _, n := range graph[vert] {
				dfs(n)
			}
		}

		path = path[:len(path)-1] // remove last element
	}

	dfs(source)
	return paths
}
