package piscine

import (
	"os"
	
	"github.com/01-edu/z01"
)

func Rot14(s string) string {
	var ans []byte
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			// Rotate lowercase letters
			ch = ((ch - 'a' + 14) % 26) + 'a'
		} else if ch >= 'A' && ch <= 'Z' {
			// Rotate uppercase letters
			ch = ((ch - 'A' + 14) % 26) + 'A'
		}
		ans = append(ans, byte(ch))
	}
	return string(ans)
}
