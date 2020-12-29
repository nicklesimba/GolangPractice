package l33tc0d3

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	movdir := 3 // up = 0, down = 1, left = 2, right = 3
	// keep t, b, l, r variables to limit boundaries
	t, b, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1

	i, j := t, l
	madeMove := false // must be set to true by loop by a case otherwise we finish
	for {
		madeMove = false
		switch movdir {
		case 0: // up
			for i = b; i >= t; i-- {
				madeMove = true
				fmt.Println(matrix[i][j])
				result = append(result, matrix[i][j])
			}
			i++        // move i back in bounds
			l++        // update left
			movdir = 3 // update movedir
		case 1: // down
			for i = t; i <= b; i++ {
				madeMove = true
				fmt.Println(matrix[i][j])
				result = append(result, matrix[i][j])
			}
			i--        // move i back in bounds
			r--        // update right
			movdir = 2 // update movedir
		case 2: // left
			for j = r; j >= l; j-- {
				madeMove = true
				fmt.Println(matrix[i][j])
				result = append(result, matrix[i][j])
			}
			j++        // move i back in bounds
			b--        // update bottom
			movdir = 0 // update movedir
		case 3: // right
			for j = l; j <= r; j++ {
				madeMove = true
				fmt.Println(matrix[i][j])
				result = append(result, matrix[i][j])
			}
			j--        // move j back in bounds
			t++        // update top
			movdir = 1 // update movedir
		}
		// end of loop iter: update i and j in whatever way possible, and update direction
		// if not possible: break from loop
		if madeMove == false {
			break
		}
	}

	return result
}
