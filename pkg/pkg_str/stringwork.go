package StringWork

import "strings"

func CaseNumbers(fileName string) int {
	/*
		-1 if no numbers
		 0 if some numbers
		 1 if all numbers
	*/
	strLen := 0
	digitCnt := 0
	for _, ch := range fileName {
		if ch == '.' {
			break
		}
		// _, err := strings.Atoi(ch) 
		// if err != nil{
		// 	digitCnt ++
		// }
		if '0' <= ch && ch <= '9' { // atoi
			digitCnt++
		}
		strLen++
	}
	if digitCnt == 0 {
		return -1
	}
	if digitCnt == strLen {
		return 1
	}
	return 0
}

func RemoveNumbers(str string) string {
	strNoNums := ""
	for _, ch := range str {
		if !('0' <= ch && ch <= '9') {
			strNoNums += string(ch)
		}
	}
	strNoNums += "\n"
	return strNoNums
}

func findDot(str string) int {
	for ind, ch := range str {
		if ch == '.' {
			return ind
		}
	}
	return -1
}

func FileNameUpper(str string) string {
	strUpper := ""
	dotPos := findDot(str)
	for i, ch := range str {
		if i < dotPos {
			strUpper += strings.ToUpper(string(ch))
		} else {
			strUpper += string(ch)
		}
	}
	strUpper += "\n"
	return strUpper
}