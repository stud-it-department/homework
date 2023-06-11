package fileUtils

import (
	"fmt"
	"github.com/goriiin/mypack/pkg/strUtils"
	"os"
)

func WalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	} else {
		file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return nil
		}
		fmt.Fprintf(file, "{path to file: %-20s\t old name: %-20s\t new name: %s}\n", path, info.Name(), strUtils.NewName(info.Name()))
		fmt.Println("Данные записаны в файл")
		return nil
	}
}
