package str

func caseNumbers(fileName string) int {
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

		if '0' <= ch && ch <= '9'{
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