package StringWork

import "strings"

func RemoveDigits(name string) string {
	s := strings.Clone(name)
	nums := "123456789"

	for _, ch := range nums {
		s = strings.ReplaceAll(s, string(ch), "")
	}

	if len(s) == 4 {
		return name + " ну такой вот он, не осуждаем, сами молодыми были"
	} else if len(s) == len(name) {
		return strings.ToUpper(name)
	}
	return s
}
