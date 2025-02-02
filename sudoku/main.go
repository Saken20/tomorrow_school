package main

import (
	"fmt"
	"os"
)

func isValid(board [9][9]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

func solve(board *[9][9]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValid(*board, row, col, num) {
						board[row][col] = num
						if solve(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func printBoard(board [9][9]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func parseInput(args []string) ([9][9]byte, error) {
	var board [9][9]byte
	if len(args) != 9 {
		return board, fmt.Errorf("Invalid input: must provide 9 rows")
	}
	for i, row := range args {
		if len(row) != 9 {
			return board, fmt.Errorf("Invalid row length: must have 9 characters")
		}
		for j, c := range row {
			if c != '.' && (c < '1' || c > '9') {
				return board, fmt.Errorf("Invalid character: %c", c)
			}
			board[i][j] = byte(c)
		}
	}
	return board, nil
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error: provide exactly 9 rows of Sudoku")
		return
	}
	board, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if solve(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error: no solution found")
	}
}
