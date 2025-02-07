package main

import (
	"os"
	"github.com/01-edu/z01"
)

func clearstr(s string) {
	var result []rune
	inSpace := false

	for _, r := range s {
		// Если текущий символ пробел
		if r == ' ' {
			// Если предыдущий символ не был пробелом, добавляем пробел в результат
			if !inSpace && len(result) > 0 {
				result = append(result, r)
			}
			inSpace = true
		} else {
			// Если текущий символ не пробел, добавляем его в результат
			result = append(result, r)
			inSpace = false
		}
	}

	// Если последним символом был пробел, удаляем его
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	// Выводим результат с использованием z01.PrintRune
	for _, r := range result {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	clearstr(os.Args[1])
}
