package piscine

import (
	"github.com/01-edu/z01"
)

const n = 8

func EightQueens() {
	solutions := make([][]int, 0)
	board := [n][n]bool{}
	placeQueens(0, board, &solutions)
	printSolutions(solutions)
}

func placeQueens(col int, board [n][n]bool, solutions *[][]int) {
	if col == n {
		saveSolution(board, solutions)
		return
	}

	for row := 0; row < n; row++ {
		if isSafe(row, col, board) {
			board[row][col] = true
			placeQueens(col+1, board, solutions)
			board[row][col] = false
		}
	}
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

func printSolutions() {
	for _, solution := range solutions {
		for _, p := range solution {
			z01.PrintRune(p)
		}
		z01.PrintRune()
	}
}

func main() {
	EightQueens()
}
