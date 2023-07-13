package pkg

import (
	"log"
	"os"
)

func RecieveDirectoryContent(path string) string {
	stack := make([]string, 0)
	stack = append(stack, path)
	names := ""
	for len(stack) > 0 {
		currentPath := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		entries, err := os.ReadDir(currentPath)

		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				newPath := currentPath + "/" + entry.Name()
				stack = append(stack, newPath)
			} else {
				names += RemoveDigits(entry.Name()) + "\n"
			}
		}
	}
	return names
}
