package StringWork

import "strings"

var nums = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func RemoveDigits(name string) string {
	s := strings.Clone(name)

	for _, ch := range nums {
		s = strings.ReplaceAll(s, string(ch), "")
	}

	if s[0] == '.' {
		return name + " ну такой вот он, не осуждаем, сами молодыми были"
	} else if len(s) == len(name) {
		return strings.ToUpper(name)
	}
	return s
}
