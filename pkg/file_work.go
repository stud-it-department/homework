package FileWork

import (
	"fmt"
	"os"
	stringWork "github.com/just-rudy/stud-hw/pkg/pkg_str"
)


func GetFileNames(dirName string) (string, string) {
	stringFileNames := ""
	file_num := 0
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println(err)
		return "", "Error in opening directory"
	} else {
		files, err := dir.ReadDir(0)
		if err != nil {
			fmt.Println(err)
			return "", "Error in going through directory"
		} else {
			for _, file := range files {
				if !file.IsDir() {
					fileName := file.Name()
					switch stringWork.CaseNumbers(fileName) {
						case -1: // upper fileName
							file_num += 1
							stringFileNames += stringWork.FileNameUpper(fileName) + "\n"
						case 0: // remove numbers from fileName
							file_num += 1
							stringFileNames += stringWork.RemoveNumbers(fileName) + "\n"
						case 1: // make a note abt it
							file_num += 1
							message := "this file " + fileName + " is special, don't touch it\n"
							stringFileNames += message
						}
				}
			}
		}
	}
	return stringFileNames, ""
}