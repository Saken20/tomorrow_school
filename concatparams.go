package piscine

func ConcatParams(args []string) string {
	text := ""
	for i := 0; i < len(args); i++ {
		text += args[i]
		if i != len(args)-1 {
			text += "\n"
		}
	}
	return text
}
