package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var queens [8]int
	solveQueens(0, &queens)
}

func solveQueens(col int, queens *[8]int) {
	if col == 8 {
		printQueens(queens)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(col, row, queens) {
			queens[col] = row
			solveQueens(col+1, queens)
		}
	}
}

func isSafe(col int, row int, queens *[8]int) bool {
	for previousCol := 0; previousCol < col; previousCol++ {
		previousRow := queens[previousCol]

		if previousRow == row {
			return false
		}

		if previousRow-row == previousCol-col || previousRow-row == col-previousCol {
			return false
		}
	}

	return true
}

func printQueens(queens *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(queens[i] + '0'))
	}
	z01.PrintRune('\n')
}
