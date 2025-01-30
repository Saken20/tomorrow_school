package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Help() {
	var s string = "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument."
	Print(s)
}

func Insert(str1 string, str2 string) string {
	return str1 + str2
}

func Order(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func Print(ans string) {
	for _, char := range ans {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		Help()
		return
	}
	var insert string
	var order bool
	var input string
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if len(arg) >= 8 && arg[:8] == "--insert" {
			if len(arg) > 8 {
				insert = arg[9:]
			}
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		}
		if arg == "--order" || arg == "-o" {
			order = true
		}
		if i == len(os.Args)-1 {
			input = arg
		}
	}
	if insert != "" {
		input = Insert(input, insert)
	}
	if order {
		input = Order(input)
	}
	Print(input)
}
