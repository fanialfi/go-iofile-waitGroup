package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fanialfi/go-iofile-waitGroup/lib"
)

// func main() {
// 	start := time.Now()
// 	var (
// 		countFile      = 100
// 		lengthFileName = 5
// 	)

// 	nameFiles := lib.CreateNameFiles(countFile, lengthFileName)
// 	lib.CreateFiles(nameFiles)

// 	finish := time.Since(start)
// 	fmt.Println(finish.String())
// }

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	files, err := os.ReadDir(lib.Path)
	if err != nil {
		log.Println(err.Error())
	}

	for _, file := range files {
		wg.Add(1)
		lib.ReadFile(file, &wg)
	}

	wg.Wait()

	finish := time.Since(start)
	fmt.Println(finish.String())
}
