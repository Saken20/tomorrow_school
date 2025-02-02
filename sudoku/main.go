package main

import (
	"fmt"
	"os"
)

type SudokuGrid [9][9]int

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var grid SudokuGrid
	for i := 1; i < 10; i++ {
		row := os.Args[i]
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}
		for j, char := range row {
			if char == '.' {
				grid[i-1][j] = 0
			} else if '1' <= char && char <= '9' {
				grid[i-1][j] = int(char - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if !isValidSudoku(&grid) {
		fmt.Println("Error")
		return
	}

	if solveSudoku(&grid) {
		printGrid(&grid)
	} else {
		fmt.Println("Error")
	}
}

func solveSudoku(grid *SudokuGrid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(grid, i, j, num) {
						grid[i][j] = num
						if solveSudoku(grid) {
							return true
						} else {
							grid[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(grid *SudokuGrid, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func isValidSudoku(grid *SudokuGrid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] != 0 {
				num := grid[i][j]
				grid[i][j] = 0
				if !isValid(grid, i, j, num) {
					return false
				}
				grid[i][j] = num
			}
		}
	}
	return true
}

func printGrid(grid *SudokuGrid) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j != 8 {
				fmt.Printf("%d ", grid[i][j])
			} else {
				fmt.Printf("%d", grid[i][j])
			}
		}
		fmt.Println()
	}
}
