package main

import "os"

const (
	MaxInt64 = 9223372036854775807
	MinInt64 = -9223372036854775808
)

func parseInt(s string) (int64, bool) {
	sign, i, num := int64(1), 0, int64(0)
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')
		if (sign == 1 && num > (MaxInt64-digit)/10) || (sign == -1 && num > -(MinInt64+digit)/10) {
			return 0, false
		}
		num = num*10 + digit
	}
	return sign * num, true
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	if n == MinInt64 {
		return "-9223372036854775808"
	}
	sign, num, res := "", n, []byte{}
	if n < 0 {
		sign, num = "-", -n
	}
	for num > 0 {
		res = append([]byte{byte('0' + num%10)}, res...)
		num /= 10
	}
	return sign + string(res)
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	a, ok1 := parseInt(args[0])
	b, ok2 := parseInt(args[2])
	if !ok1 || !ok2 {
		return
	}
	switch args[1] {
	case "+":
		if res, ok := a+b, (b > 0 && a > MaxInt64-b) || (b < 0 && a < MinInt64-b); ok {
			return
		} else {
			os.Stdout.WriteString(itoa(res) + "\n")
		}
	case "-":
		if res, ok := a-b, (b < 0 && a > MaxInt64+b) || (b > 0 && a < MinInt64+b); ok {
			return
		} else {
			os.Stdout.WriteString(itoa(res) + "\n")
		}
	case "*":
		if (a == MinInt64 && b == -1) || (b != 0 && a > MaxInt64/b) || (b != 0 && a < MinInt64/b) {
			return
		}
		os.Stdout.WriteString(itoa(a*b) + "\n")
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		if a == MinInt64 && b == -1 {
			return
		}
		os.Stdout.WriteString(itoa(a/b) + "\n")
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		os.Stdout.WriteString(itoa(a%b) + "\n")
	}
}
