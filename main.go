package main

/*
Если цифры в названии - убрать числа
Если только цифры     - сделать пометку
Если нет цифр         - upper
*/

import (
	// pkg "./pkg"
	"fmt"
	"os"
	"strings"
)

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

		if '0' <= ch && ch <= '9' {
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

func removeNumbers(str string) string {
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

func fileNameUpper(str string) string {
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

func main() {
	var dirName string
	var infoFile string

	// input output file
	fmt.Printf("input info-file title: ")
	fmt.Scanf("%s", &infoFile)

	// input directory name
	fmt.Printf("input dir name: ")
	fmt.Scanf("%s", &dirName)

	// open dir and read files in it
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println(err)
	} else {
		files, err := dir.ReadDir(0)
		if err != nil {
			fmt.Println(err)
		} else {
			// open output file for writing
			outFile, err := os.Create(infoFile)
			if err != nil {
				fmt.Println(err)
			} else {
				for _, file := range files {
					if !file.IsDir() {
						fileName := file.Name()
						switch caseNumbers(fileName) {
						case -1: // upper fileName
							outFile.WriteString(fileNameUpper(fileName))
							//fmt.Println(fileNameUpper(fileName))
						case 0: // remove numbers from fileName
							outFile.WriteString(removeNumbers(fileName))
							// fmt.Println(removeNumbers(fileName))
						case 1: // make a note abt it
							message := "this file " + fileName + " is special, don't touch it\n"
							outFile.WriteString(message)
						}
					}
				}
				outFile.Close()
			}
		}
	}
}
