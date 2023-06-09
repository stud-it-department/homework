package Files

import (
	s "./StringWork"
	"log"
	"os"
)

func GetDirectoryContent(path string, names *string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		//fmt.Println(entry.Name())
		if entry.IsDir() {
			newPath := path + "/" + entry.Name()
			GetDirectoryContent(newPath, names)
		} else {
			*names += s.RemoveDigits(entry.Name()) + "\n"
		}

	}
}
