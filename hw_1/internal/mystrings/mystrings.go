package mystrings

import (
	"strconv"
	"strings"
)

func NameType(name string) int {
	nums := 0
	length := len(name)
	for _, symb := range name {
		_, err := strconv.Atoi(string(symb))
		if err == nil {
			nums += 1
		}
	}

	if nums == length {
		return 2
	} else if nums > 0 {
		return 1
	} else {
		return 0
	}

}
func ChooseLine(name string, val int) string {
	line := ""
	if val == 2 {
		line = "Этот " + name + " особенный"
	} else if val == 1 {
		line = DeleteNums(name)
	} else if val == 0 {
		line = strings.ToUpper(name)
	}
	return line
}

func DeleteNums(name string) string {
	line := ""
	for _, symb := range name {
		_, err := strconv.Atoi(string(symb))
		if err != nil {
			line = line + string(symb)
		}
	}
	return line
}
