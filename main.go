package main

/*
Если цифры в названии - убрать числа
Если только цифры     - сделать пометку
Если нет цифр         - upper
*/

import (
	"fmt"
	"os"
	s "github.com/just-rudy/stud-hw/pkg"
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

	n := 0
	allFileNames := ""
	status := s.GetFileNames(dirName, &allFileNames, &n)
	if status == 0 {
		outFile, err := os.Create(infoFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile.WriteString(allFileNames)
		outFile.Close()
	}	
}
