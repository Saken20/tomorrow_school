package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	QuadChecker()
	fmt.Println()
}

func QuadChecker() {
	var x int
	y := 0
	var count int
	stdin, _ := io.ReadAll(os.Stdin)
	input := string(stdin)
	for i, v := range input {
		if v == '\n' {
			y += 1
		}
		if y == 0 {
			x = i + 1 // input = "ABA\nB B"
		}
	}
	if input == QuadA(x, y) {
		PrintAns(x, y, &count, 'A')
	}
	if input == QuadB(x, y) {
		PrintAns(x, y, &count, 'B')
	}
	if input == QuadC(x, y) {
		PrintAns(x, y, &count, 'C')
	}
	if input == QuadD(x, y) {
		PrintAns(x, y, &count, 'D')
	}
	if input == QuadE(x, y) {
		PrintAns(x, y, &count, 'E')
	}
	if count == 0 {
		fmt.Print("Not a Quad function")
		return
	}
}

func PrintAns(x, y int, count *int, r rune) {
	*count++
	if *count > 1 {
		fmt.Print(" || ")
	}
	fmt.Printf("[quad%s] [%d] [%d]", string(r), x, y)
}

func GetNumber() (int, int, error) {
	args := os.Args[1:]
	x, err := strconv.Atoi(args[0])
	y, err := strconv.Atoi(args[1])

	return x, y, err
}

func QuadA(x, y int) string {
	out := ""
	if x < 1 || y < 1 {
		return ""
	}
	if y >= 1 {
		out = out + "o"
		for i := 0; i < (x - 2); i++ {
			out = out + "-"
		}
		if x > 1 {
			out = out + "-"
		}
	}
	if y > 2 {
		for i := 0; i < (y - 2); i++ {
			out += "\n"
			out += "|"
			if x == 2 {
				out += "|"
			}
			if x > 2 {
				for i := 0; i < (x - 2); i++ {
					out += " "
				}
				out += "|"
			}
		}
	}
	if y > 1 {
		out += "\n"
		if x >= 1 {
			out += "o"
		}
		for i := 0; i < (x - 2); i++ {
			out += "-"
		}
		if x > 1 {
			out += "o"
		}
	}
	out += "\n"
	return out
}

func QuadB(x, y int) string {
	out := ""
	if x < 1 || y < 1 {
		return ""
	}
	if y >= 1 {
		out += "/"
		for i := 0; i < (x - 2); i++ {
			out += "*"
		}
		if x > 1 {
			out += "\\"
		}
	}
	if y > 2 {
		for i := 0; i < (y - 2); i++ {
			out += "\n"
			out += "*"
			if x == 2 {
				out += "*"
			}
			if x > 2 {
				for i := 0; i < (x - 2); i++ {
					out += " "
				}
				out += "*"
			}
		}
	}
	if y > 1 {
		out += "\n"
		if x >= 1 {
			out += "\\"
		}
		for i := 0; i < (x - 2); i++ {
			out += "*"
		}
		if x > 1 {
			out += "/"
		}
	}
	out += "\n"
	return out
}

func QuadC(x, y int) string {
	out := ""
	if x < 1 || y < 1 {
		return ""
	}
	if y >= 1 {
		out += "A"
		for i := 0; i < (x - 2); i++ {
			out += "B"
		}
		if x > 1 {
			out += "A"
		}
	}
	if y > 2 {
		for i := 0; i < (y - 2); i++ {
			out += "\n"
			out += "B"
			if x == 2 {
				out += "B"
			}
			if x > 2 {
				for i := 0; i < (x - 2); i++ {
					out += " "
				}
				out += "B"
			}
		}
	}
	if y > 1 {
		out += "\n"
		if x >= 1 {
			out += "C"
		}
		for i := 0; i < (x - 2); i++ {
			out += "B"
		}
		if x > 1 {
			out += "C"
		}
	}
	out += "\n"
	return out
}

func QuadD(x, y int) string {
	out := ""
	if x < 1 || y < 1 {
		return ""
	}
	if y >= 1 {
		out += "A"
		for i := 0; i < (x - 2); i++ {
			out += "B"
		}
		if x > 1 {
			out += "C"
		}
	}
	if y > 2 {
		for i := 0; i < (y - 2); i++ {
			out += "\n"
			out += "B"
			if x == 2 {
				out += "B"
			}
			if x > 2 {
				for i := 0; i < (x - 2); i++ {
					out += " "
				}
				out += "B"
			}
		}
	}
	if y > 1 {
		out += "\n"
		if x >= 1 {
			out += "A"
		}
		for i := 0; i < (x - 2); i++ {
			out += "B"
		}
		if x > 1 {
			out += "C"
		}
	}
	out += "\n"
	return out
}

func QuadE(x, y int) string {
	out := ""
	if x < 1 || y < 1 {
		return ""
	}
	if y >= 1 {
		out += "A"
		for i := 0; i < (x - 2); i++ {
			out += "B"
		}
		if x > 1 {
			out += "C"
		}
	}
	if y > 2 {
		for i := 0; i < (y - 2); i++ {
			out += "\n"
			out += "B"
			if x == 2 {
				out += "B"
			}
			if x > 2 {
				for i := 0; i < (x - 2); i++ {
					out += " "
				}
				out += "B"
			}
		}
	}
	if y > 1 {
		out += "\n"
		if x >= 1 {
			out += "C"
		}
		for i := 0; i < (x - 2); i++ {
			out += "B"
		}
		if x > 1 {
			out += "A"
		}
	}
	out += "\n"
	return out
}
