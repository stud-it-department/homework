package myfiles

import "os"

func CreateFile(name string) *os.File {
	file, _ := os.Create(name)
	//defer file.Close()
	return file
}

func CloseFile(file *os.File) {
	file.Close()
}

func WriteLine(file *os.File, line string) {
	_, _ = file.WriteString(line + "\n")
}
