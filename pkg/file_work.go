package FileWork

import (
	"fmt"
	"os"
	s "github.com/just-rudy/stud-hw/pkg/pkg_str"
)


func GetFileNames(dirName string, sliceFileNames *[]string, n *int) int {
	*n = 0
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println(err)
		return 1
	} else {
		files, err := dir.ReadDir(0)
		if err != nil {
			fmt.Println(err)
			return 1
		} else {
			for _, file := range files {
				if !file.IsDir() {
					fileName := file.Name()
					switch s.CaseNumbers(fileName) {
						case -1: // upper fileName
							*n += 1
							*sliceFileNames = append(*sliceFileNames, s.FileNameUpper(fileName))
						case 0: // remove numbers from fileName
							*n += 1
							*sliceFileNames = append(*sliceFileNames, s.RemoveNumbers(fileName))
						case 1: // make a note abt it
							message := "this file " + fileName + " is special, don't touch it\n"
							*n += 1
							*sliceFileNames = append(*sliceFileNames, message)
						}
				}
			}
		}
	}
	return 0
}