package main

import (
	"log"
	"os"
	"time"
)

func main()  {
	touch()
}

func touch() {
	if len(os.Args) < 2 {
		log.Fatal("You need to provide a file name.")
	}
	fileName := os.Args[1]
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		createFile(fileName)
	} else {
		changeFileTime(fileName)
	}
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func changeFileTime(fileName string) {
	currentTime := time.Now().Local()
	err := os.Chtimes(fileName, currentTime, currentTime)
	if err != nil {
		log.Fatal(err)
	}
}