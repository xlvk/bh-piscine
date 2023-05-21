package piscine

import (
	"github.com/01-edu/z01"
)

const n = 8

var (
	board     [n][n]bool
	solutions [][]int
)

func EightQueens() {
	solutions = [][]int{}
	placeQueens(0)
	printSolutions()
}

func placeQueens(col int) {
	if col == n {
		saveSolution()
		return
	}

	for row := 0; row < n; row++ {
		if isSafe(row, col) {
			board[row][col] = true
			placeQueens(col + 1)
			board[row][col] = false
		}
	}
}

func isSafe(row, col int) bool {
	// Check row and column
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

func saveSolution() {
	solution := []int{}
	for col := 0; col < n; col++ {
		for row := 0; row < n; row++ {
			if board[row][col] {
				solution = append(solution, row+1)
			}
		}
	}
	solutions = append(solutions, solution)
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
