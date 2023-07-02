package main

/*
Если цифры в названии - убрать числа
Если только цифры     - сделать пометку
Если нет цифр         - upper
*/

import (
	"fmt"
	"os"
	s "github.com/just-rudy/stud-hw/StringWork"

)



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
						switch s.CaseNumbers(fileName) {
						case -1: // upper fileName
							outFile.WriteString(s.FileNameUpper(fileName))
						case 0: // remove numbers from fileName
							outFile.WriteString(s.RemoveNumbers(fileName))
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
