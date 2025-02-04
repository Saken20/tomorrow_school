package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printContent(file *os.File) {
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			// Печатаем каждый байт как руну
			for _, b := range buf[:n] {
				z01.PrintRune(rune(b))
			}
		}
		if err == io.EOF {
			break // Конец файла
		}
		if err != nil {
			break // Ошибка при чтении
		}
	}
}

func main() {
	// Если файлов нет, читаем из stdin
	if len(os.Args) < 2 {
		printContent(os.Stdin)
		return
	}
	// Обрабатываем каждый файл
	for _, filePath := range os.Args[1:] {
		file, err := os.Open(filePath)
		if err != nil {
			// Печатаем ошибку и завершаем с кодом 1
			for _, char := range "ERROR: " + err.Error() + "\n" {
				z01.PrintRune(char)
			}
			os.Exit(1)
		}
		printContent(file)
		file.Close()
	}
}
