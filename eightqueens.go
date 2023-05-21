package piscine

import (
	"github.com/01-edu/z01"
)

const n = 8

func EightQueens() {
	solutions := make([][]int, 0)
	board := [n][n]bool{}
	col := 0

	var placeQueens func(col int, board [n][n]bool)
	placeQueens = func(col int, board [n][n]bool) {
		if col == n {
			saveSolution(board, &solutions)
			return
		}

		for row := 0; row < n; row++ {
			if isSafe(row, col, board) {
				board[row][col] = true
				placeQueens(col+1, board)
				board[row][col] = false
			}
		}
	}

	placeQueens(col, board)
	printSolutions(solutions)
	z01.Flush()
}

func isSafe(row, col int, board [n][n]bool) bool {
	// Check row
	for i := 0; i < col; i++ {
		if board[row][i] {
			return false
		}
	}

	// Check upper diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// Check lower diagonal
	for i, j := row+1, col-1; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

func saveSolution(board [n][n]bool, solutions *[][]int) {
	solution := make([]int, 0)
	for col := 0; col < n; col++ {
		for row := 0; row < n; row++ {
			if board[row][col] {
				solution = append(solution, row+1)
			}
		}
	}
	*solutions = append(*solutions, solution)
}

func printSolutions(solutions [][]int) {
	for _, solution := range solutions {
		for _, pos := range solution {
			if pos >= 1 && pos <= 8 {
				z01.PrintRune(rune(pos + '0'))
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	EightQueens()
}
