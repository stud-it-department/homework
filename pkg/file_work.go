package FileWork

import (
	"fmt"
	"os"
	s "github.com/just-rudy/stud-hw/pkg/pkg_str"
)


func GetFileNames(dirName string, stringFileNames *string, n *int) int {
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
							*stringFileNames += s.FileNameUpper(fileName) + "\n"
						case 0: // remove numbers from fileName
							*n += 1
							*stringFileNames += s.RemoveNumbers(fileName) + "\n"
						case 1: // make a note abt it
							*n += 1
							message := "this file " + fileName + " is special, don't touch it\n"
							*stringFileNames += message
						}
				}
			}
		}
	}
	return 0
}