package main

import (
	f "./Files"
	"fmt"
	"os"
)

func main() {

	fmt.Println("enter the path to \"names\" file")
	namesPath := ""
	fmt.Scanln(&namesPath)

	fd, err := os.OpenFile(namesPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("enter the path to root")
	rootPath := ""
	fmt.Scanln(&rootPath)

	names := ""
	f.GetDirectoryContent(rootPath, &names)
	//fmt.Println(names)

	if _, err := fd.WriteString(names); err != nil {
		fmt.Println(err)
		return
	}
}

// /home/kostya/StudIT/Test/names.txt
// /home/kostya/StudIT/Test
