package strUtils

import (
	"strings"
	"unicode"
)

func NewName(name string) string {

	arr := strings.Split(name, ".")
	nameWithoutNums := ""
	var flag bool
	for _, char := range arr[0] {
		if !unicode.IsDigit(char) {
			nameWithoutNums += string(char)
		} else {
			flag = true
		}
	}
	name = nameWithoutNums + "." + arr[1]
	if !flag {
		name = strings.ToUpper(name)
	}
	return name
}

//func toUpperCase(name *string) string {
//
//	ext := filepath.Ext(*name)
//	filename := strings.TrimSuffix(*name, ext)
//	return strings.ToUpper(filename) + ext
//
//}
