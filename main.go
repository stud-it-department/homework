package main

import (
	"bufio"
	"fmt"
	"github.com/goriiin/mypack/pkg/fileUtils"
	"os"
	"path/filepath"
)

func main() {
	out, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка с файлом вывода:", err)
		return
	}
	out.Close()

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		root := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка чтения файла:", err)
			return
		} else {

			if err := filepath.Walk(root, fileUtils.WalkFunc); err != nil {
				fmt.Printf("Какая-то ошибка: %v\n", err)
			}
		}
	}

}
