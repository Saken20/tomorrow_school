package piscine

import "github.com/01-edu/z01"

func Rot14(s string) string {
	ans := ""
	for _, ch := range s {
		ans = append(ans, string(ch + 14))
	}
	return ans
}
