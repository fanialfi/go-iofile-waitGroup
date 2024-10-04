package lib

import (
	"fmt"
	"log"
	"os"
)

var Path = "/tmp/golang"

func CreateNameFiles(count, lengthNameFile int) []string {
	var nameFiles = make([]string, count)

	for index := range nameFiles {
		nameFiles[index] = fmt.Sprintf("%s.txt", RandomString(lengthNameFile))
	}

	return nameFiles
}

func CreateFiles(nameFiles []string) {
	var pathFile = make([]string, len(nameFiles))
	for index := range pathFile {
		pathFile[index] = fmt.Sprintf("%s/%s", Path, nameFiles[index])
	}

	for _, value := range pathFile {
		_, err := os.Stat(value)
		if os.IsNotExist(err) {
			file, err := os.Create(value)
			if err != nil {
				log.Panicln(err.Error())
			}
			// defer function handle closing file
			defer func() {
				err := file.Close()

				if err != nil {
					log.Panicln(err.Error())
				}
			}()

			// write string sebanyak 5MB
			_, err = file.WriteString(RandomString(5 << 20))
			if err != nil {
				log.Panicln(err.Error())
			}
		}
	}
}
