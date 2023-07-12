package main

import (
	"fmt"
	"os"
	"try/internal/myfiles"
	"try/internal/mystrings"
)

func GetFileNames(path string, names *[]string) {
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, file := range files {
		*names = append(*names, file.Name())
		fmt.Println(file.Name())
	}

}

func main() {
	var path string
	fmt.Scanf("%s", &path)

	names := []string{}
	GetFileNames(path, &names)

	file := myfiles.CreateFile("output.txt")

	for _, name := range names {
		v := mystrings.NameType(name)
		line := mystrings.ChooseLine(name, v)
		myfiles.WriteLine(file, line)
	}

	myfiles.CloseFile(file)

}
