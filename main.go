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
	var sliceFileNames []string

	// input output file
	fmt.Printf("input info-file title: ")
	fmt.Scanf("%s", &infoFile)

	// input directory name
	fmt.Printf("input dir name: ")
	fmt.Scanf("%s", &dirName)

	n := 0
	status := s.GetFileNames(dirName, &sliceFileNames, &n)
	if status == 0 {
		outFile, err := os.Create(infoFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i:= 0; i < n; i++ {
			outFile.WriteString(sliceFileNames[i])
		outFile.Close()
		}
	}	
}
