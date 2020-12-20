package l33tc0d3

func findPaths(m int, n int, N int, i int, j int) int {
	// fmt.Println(m, n, N, i, j)

	// Naive approach
	// return findPathsNaive(m, n, N, i, j)

	// Memoized approach
	pathsArr := make([][][]int, m)
	for i := range pathsArr {
		pathsArr[i] = make([][]int, n)
		for j := range pathsArr[i] {
			pathsArr[i][j] = make([]int, N+1)
			for k := 0; k < N+1; k++ {
				pathsArr[i][j][k] = -1
			}
		}
	}
	return findPathsMemo(m, n, N, i, j, &pathsArr)

	// Dynamic Programming Approach
	// return findPathsDP(m, n, N, i, j)
}

// Naive because it calculates situations redundantly instead of cacheing previously computed values.
func findPathsNaive(m int, n int, N int, i int, j int) int {
	// Base case, when you're out of bounds
	if i < 0 || i >= m || j < 0 || j >= n {
		return 1
	} else if N == 0 {
		return 0
	}

	// Recursive case, 1 or more moves remain, check up/down/left/right with N-1 moves
	return findPathsNaive(m, n, N-1, i, j-1) + findPathsNaive(m, n, N-1, i, j+1) + findPathsNaive(m, n, N-1, i-1, j) + findPathsNaive(m, n, N-1, i+1, j)
}

// Second try using memoization
func findPathsMemo(m int, n int, N int, i int, j int, pathsArr *[][][]int) int {
	// Base case, when you're out of bounds
	if i < 0 || i >= m || j < 0 || j >= n {
		return 1
	} else if N == 0 {
		return 0
	}

	if (*pathsArr)[i][j][N] >= 0 {
		// fmt.Println("val[", i, "][", j, "][", N, "]:", (*pathsArr)[i][j][N])
		return (*pathsArr)[i][j][N]
	}

	// Recursive case, 1 or more moves remain, check up/down/left/right with N-1 moves
	(*pathsArr)[i][j][N] = (findPathsMemo(m, n, N-1, i, j-1, pathsArr) + findPathsMemo(m, n, N-1, i, j+1, pathsArr) + findPathsMemo(m, n, N-1, i-1, j, pathsArr) + findPathsMemo(m, n, N-1, i+1, j, pathsArr)) % 1000000007
	return (*pathsArr)[i][j][N]
}

// Scratch work
/*
Memo idea
Keep a 3d array, i/j for coordinate, k for # of moves left.
Run recursive algorithm, and record a position/# of moves left with 0 or 1 for no path to out of bounds or path to out of bounds.
To make use of the memo array, return values from it for indices that are non-negative!
*/

/*
DP idea
keep a 2D DP array, increment a position for every time it has a way to get out of bounds
start from starting position
move up/down/left/right and store
*/
