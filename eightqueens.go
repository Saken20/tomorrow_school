package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	var usedDiagonals1 [15]bool
	var usedDiagonals2 [15]bool
	var usedColumns [8]bool
	var solve func(row int)
	solve = func(row int) {
		if row == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + 1 + '0'))
			}
			z01.PrintRune('\n')
			return
		}
		for col := 0; col < 8; col++ {
			if !usedColumns[col] && !usedDiagonals1[row-col+7] && !usedDiagonals2[row+col] {
				board[row] = col
				usedColumns[col] = true
				usedDiagonals1[row-col+7] = true
				usedDiagonals2[row+col] = true
				solve(row + 1)
				usedColumns[col] = false
				usedDiagonals1[row-col+7] = false
				usedDiagonals2[row+col] = false
			}
		}
	}
	solve(0)
}
