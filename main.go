package main

import (
	"fmt"
	f "github.com/c9llm3bones/homework/pkg/pkg"
	"log"
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

	names := f.RecieveDirectoryContent(rootPath)
	fmt.Println(names)

	er := fd.Close()
	if er != nil {
		return
	}

	log.Fatal(err)

}

// /home/kostya/StudIT/Test/names.txt
// /home/kostya/StudIT/Test
