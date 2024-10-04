package lib

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"sync"
)

func ReadFile(dirEntry fs.DirEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	path := fmt.Sprintf("%s/%s", Path, dirEntry.Name())

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Panicln(err.Error())
	}

	// function for handle file closing
	defer func() {
		err := file.Close()

		if err != nil {
			log.Panicln(err.Error())
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Panicln(err.Error())
	}

	bytes := make([]byte, fileInfo.Size())

	n, err := file.Read(bytes)
	if err != nil && err != io.EOF {
		log.Panicln(err.Error())
	}
	if n <= 0 {
		return
	}
	fmt.Println(string(bytes))
}
